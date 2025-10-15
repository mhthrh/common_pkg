package pubsub

import (
	"context"
	"fmt"
	"strings"

	cfg "github.com/mhthrh/common_pkg/pkg/model/config"
	"github.com/mhthrh/common_pkg/pkg/xSolace"
	"solace.dev/go/messaging"
	"solace.dev/go/messaging/pkg/solace"
	"solace.dev/go/messaging/pkg/solace/config"
	"solace.dev/go/messaging/pkg/solace/message"
	"solace.dev/go/messaging/pkg/solace/resource"
)

type SolacePubSub struct {
	messagingService solace.MessagingService
}

func replaceLastComma(s, replacement string) string {
	index := strings.LastIndex(s, ",")
	if index == -1 {
		return s // no comma found
	}
	return s[:index] + replacement + s[index+1:]
}
func New(cfg cfg.Solace) (*SolacePubSub, error) {
	host := ""
	for _, v := range cfg.Hosts {
		host += fmt.Sprintf("%s://%s:%d,", v.Schema, v.Host, v.Port)
	}

	brokerConfig := config.ServicePropertyMap{
		config.TransportLayerPropertyHost:                replaceLastComma(host, ""),
		config.ServicePropertyVPNName:                    cfg.VPN,
		config.AuthenticationPropertySchemeBasicPassword: cfg.Pass,
		config.AuthenticationPropertySchemeBasicUserName: cfg.User,
	}
	messagingService, err := messaging.NewMessagingServiceBuilder().FromConfigurationProvider(brokerConfig).Build()
	if err != nil {
		return nil, err
	}
	if messagingService.Connect() != nil {
		return nil, err
	}
	fmt.Println("Connected to the broker? ", messagingService.IsConnected())

	return &SolacePubSub{
		messagingService: messagingService,
	}, nil

}

func (s *SolacePubSub) Publish(ctx context.Context, p *xSolace.Pipe) {
	persistentPublisher, builderErr := s.messagingService.CreatePersistentMessagePublisherBuilder().Build()
	if builderErr != nil {
		panic(builderErr)
	}
	persistentPublisher.SetMessagePublishReceiptListener(func(receipt solace.PublishReceipt) {
		fmt.Println("Received a Publish Receipt from the broker")
		p.ReceiptOut <- "ok123"
		// fmt.Println("IsPersisted: ", receipt.IsPersisted())
		// fmt.Println("Message : ", receipt.GetMessage())
		if receipt.GetError() != nil {
			fmt.Println("Gauranteed Message is NOT persisted on the broker! Received NAK")
			fmt.Println("Error is: ", receipt.GetError())
			// probably want to do something here.  some error handling possibilities:
			//  - send the message again
			//  - send it somewhere else (error handling queue?)
			//  - log and continue
			//  - pause and retry (backoff) - maybe set a flag to slow down the publisher
		}

	})

	startErr := persistentPublisher.Start()
	if startErr != nil {
		panic(startErr)
	}

	fmt.Println("Persistent Publisher running? ", persistentPublisher.IsRunning())

	topic := resource.TopicOf(p.Topic)
	fmt.Printf("Publishing on: %s, please ensure queue has matching subscription.\n", topic.GetName())
	messageBuilder := s.messagingService.MessageBuilder().
		WithProperty("application", "samples").
		WithProperty("language", "go")
	go func() {
		for persistentPublisher.IsReady() {
			select {
			case <-ctx.Done():
				return
			case msg := <-p.MsgIn:
				message, err := messageBuilder.BuildWithStringPayload(msg)
				if err != nil {
					panic(err)
				}

				publishErr := persistentPublisher.Publish(message, topic, nil, nil)

				if publishErr != nil {
					panic(publishErr)
				}
			}
		}
	}()

}

func (s *SolacePubSub) Listen(ctx context.Context, p *xSolace.Pipe) {
	// queueName := "durable-queue"
	// durableExclusiveQueue := resource.QueueDurableExclusive("durable-queue")
	queueName := "nondurable-queue"
	nonDurableExclusiveQueue := resource.QueueNonDurableExclusive("nondurable-queue")
	topic := resource.TopicSubscriptionOf(p.Topic)

	// Build a Gauranteed message receiver and bind to the given queue
	strategy := config.MissingResourcesCreationStrategy("CREATE_ON_START")
	// Durable Queue
	persistentReceiver, err := s.messagingService.CreatePersistentMessageReceiverBuilder().WithMessageAutoAcknowledgement().WithMissingResourcesCreationStrategy(strategy).WithSubscriptions(topic).Build(nonDurableExclusiveQueue)

	if err != nil {
		panic(err)
	}
	// Non-durable Queue
	// persistentReceiver, err := messagingService.CreatePersistentMessageReceiverBuilder().WithMissingResourcesCreationStrategy(strategy).WithSubscriptions(topic).Build(nonDurableExclusiveQueue)

	// Handling a panic from a non existing queue
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Make sure queue name '%s' exists on the broker.\nThe following error occurred when attempting to connect to create a Persistent Message Receiver:\n%s", queueName, err)
		}
	}()

	// Start Persistent Message Receiver
	if err := persistentReceiver.Start(); err != nil {
		panic(err)
	}

	fmt.Println("Persistent Receiver running? ", persistentReceiver.IsRunning())

	// Register Message callback handler to the Message Receiver
	if regErr := persistentReceiver.ReceiveAsync(func(inboundMessage message.InboundMessage) {
		var messageBody string

		if payload, ok := inboundMessage.GetPayloadAsString(); ok {
			p.MsgOut <- payload
		} else if payload, ok := inboundMessage.GetPayloadAsBytes(); ok {
			p.MsgOut <- string(payload)
		}

		fmt.Printf("Received Message Body %s \n", messageBody)
		// fmt.Printf("Message Dump %s \n", message)
	}); regErr != nil {
		panic(regErr)
	}
	fmt.Printf("\n Bound to queue: %s\n", queueName)

}
