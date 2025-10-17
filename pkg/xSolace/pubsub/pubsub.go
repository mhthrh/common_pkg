package pubsub

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
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
	log.Println("Connected to the broker? ", messagingService.IsConnected())

	return &SolacePubSub{
		messagingService: messagingService,
	}, nil

}

func (s *SolacePubSub) Publish(ctx context.Context, p *xSolace.Pipe, msg, topicName string) error {

	persistentPublisher, builderErr := s.messagingService.CreatePersistentMessagePublisherBuilder().Build()
	if builderErr != nil {
		panic(builderErr)
	}
	persistentPublisher.SetMessagePublishReceiptListener(func(receipt solace.PublishReceipt) {
		p.ReceiptOut <- "ok123"

		if receipt.GetError() != nil {

		}

	})

	startErr := persistentPublisher.Start()
	if startErr != nil {
		panic(startErr)
	}

	log.Println("Persistent Publisher running? ", persistentPublisher.IsRunning())

	topic := resource.TopicOf(topicName)
	log.Printf("Publishing on: %s, please ensure queue has matching subscription.\n", topic.GetName())
	messageBuilder := s.messagingService.MessageBuilder().
		WithProperty("application", "samples").
		WithProperty("language", "go")

	if persistentPublisher.IsReady() {
		message, err := messageBuilder.BuildWithStringPayload(msg)
		if err != nil {
			panic(err)
		}

		publishErr := persistentPublisher.Publish(message, topic, nil, nil)

		if publishErr != nil {
			panic(publishErr)
		}
		return nil
	}
	return errors.New("persistent Publisher did not start")

}

func (s *SolacePubSub) Listen(ctx context.Context, p *xSolace.Pipe, topicName string) {

	queueName := "nondurable-queue"
	nonDurableExclusiveQueue := resource.QueueNonDurableExclusive("nondurable-queue")
	topic := resource.TopicSubscriptionOf(topicName)

	strategy := config.MissingResourcesCreationStrategy("CREATE_ON_START")
	persistentReceiver, err := s.messagingService.CreatePersistentMessageReceiverBuilder().WithMessageAutoAcknowledgement().WithMissingResourcesCreationStrategy(strategy).WithSubscriptions(topic).Build(nonDurableExclusiveQueue)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Make sure queue name '%s' exists on the broker.\nThe following error occurred when attempting to connect to create a Persistent Message Receiver:\n%s", queueName, err)
		}
	}()

	if err := persistentReceiver.Start(); err != nil {
		panic(err)
	}

	log.Println("Persistent Receiver running? ", persistentReceiver.IsRunning())

	if regErr := persistentReceiver.ReceiveAsync(func(inboundMessage message.InboundMessage) {

		if payload, ok := inboundMessage.GetPayloadAsString(); ok {
			p.MsgOut <- payload
		} else if payload, ok := inboundMessage.GetPayloadAsBytes(); ok {
			p.MsgOut <- string(payload)
		}

	}); regErr != nil {
		panic(regErr)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
}
