package pubsub

import (
	"errors"
	"fmt"

	cfg "github.com/mhthrh/common_pkg/pkg/model/config"
	"github.com/mhthrh/common_pkg/pkg/xSolace"
	"solace.dev/go/messaging"
	"solace.dev/go/messaging/pkg/solace"
	"solace.dev/go/messaging/pkg/solace/config"
	"solace.dev/go/messaging/pkg/solace/resource"
)

type SolacePubSub struct {
	messagingService    solace.MessagingService
	persistentPublisher solace.PersistentMessagePublisher
	directReceiver      solace.DirectMessageReceiver
}

func New(cfg cfg.Solace, msg xSolace.Message) (*SolacePubSub, error) {
	host := ""
	for _, v := range cfg.Hosts {
		host += fmt.Sprintf("%s://%s:%d,", v.Schema, v.Host, v.Port)
	}

	brokerConfig := config.ServicePropertyMap{
		config.TransportLayerPropertyHost:                host,
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

	return &SolacePubSub{
		messagingService: messagingService,
		directReceiver:   nil,
	}, nil

}

func (s *SolacePubSub) Publish(msg string, m xSolace.Message) error {
	if !s.persistentPublisher.IsReady() {
		return errors.New("persistent Publisher is not ready")
	}
	messageBuilder := s.messagingService.MessageBuilder().
		WithProperty("application", "samples").
		WithProperty("language", "go")

	message, err := messageBuilder.BuildWithStringPayload(msg)
	if err != nil {
		panic(err)
	}
	topic := resource.TopicOf(m.Subscription)

	// Publish on dynamic topic with dynamic body
	// NOTE: publishing to topic, so make sure GuaranteedReceiver queue is subscribed to same topic,
	//       or enable "Reject Message to Sender on No Subscription Match" the client-profile
	return s.persistentPublisher.Publish(message, topic, nil, nil)
	// Block until message is acknowledged
	// publishErr := persistentPublisher.PublishAwaitAcknowledgement(message, topic, 2*time.Second, nil)

}

func (s *SolacePubSub) Listen(fnc func(receipt solace.PublishReceipt)) error {
	persistentPublisher, builderErr := s.messagingService.CreatePersistentMessagePublisherBuilder().Build()
	if builderErr != nil {
		panic(builderErr)
	}

	// Set the message publisher receipt listener
	persistentPublisher.SetMessagePublishReceiptListener(fnc)
	fmt.Println("Persistent Publisher running? ", persistentPublisher.IsRunning())

	return persistentPublisher.Start()
}
