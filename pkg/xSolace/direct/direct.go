package direct

import (
	"fmt"

	cfg "github.com/mhthrh/common_pkg/pkg/model/config"
	"github.com/mhthrh/common_pkg/pkg/xSolace"
	"solace.dev/go/messaging"
	"solace.dev/go/messaging/pkg/solace"
	"solace.dev/go/messaging/pkg/solace/config"
	"solace.dev/go/messaging/pkg/solace/message"
	"solace.dev/go/messaging/pkg/solace/resource"
)

type SolaceDirect struct {
	messagingService solace.MessagingService
	directPublisher  solace.DirectMessagePublisher
	directReceiver   solace.DirectMessageReceiver
}

func New(cfg cfg.Solace, msg xSolace.Message) (*SolaceDirect, error) {
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
	fmt.Println("Connected to the broker? ", messagingService.IsConnected())

	//  Build a Direct Message Publisher
	directPublisher, builderErr := messagingService.CreateDirectMessagePublisherBuilder().Build()
	if builderErr != nil {
		return nil, err
	}

	if directPublisher.Start() != nil {
		return nil, err
	}
	fmt.Println("Direct Publisher running? ", directPublisher.IsRunning())
	directReceiver, err := messagingService.CreateDirectMessageReceiverBuilder().
		WithSubscriptions(resource.TopicSubscriptionOf(msg.Subscription)).
		Build()

	if err != nil {
		return nil, err
	}
	fmt.Println("Direct Receiver running? ", directReceiver.IsRunning())

	return &SolaceDirect{
		messagingService: messagingService,
		directPublisher:  directPublisher,
		directReceiver:   directReceiver,
	}, nil
}
func (s *SolaceDirect) Publish(msg string, m xSolace.Message) error {
	messageBuilder := s.messagingService.MessageBuilder().
		WithProperty("application", "samples").
		WithProperty("language", "go")

	message, err := messageBuilder.BuildWithStringPayload(msg)
	if err != nil {
		return err
	}
	return s.directPublisher.Publish(message, resource.TopicOf(m.Subscription))

}

func (s *SolaceDirect) Listen(fnc func(message.InboundMessage)) error {
	if err := s.directReceiver.Start(); err != nil {
		return err
	}

	fmt.Println("Direct Receiver running? ", s.directReceiver.IsRunning())

	return s.directReceiver.ReceiveAsync(fnc)
}
