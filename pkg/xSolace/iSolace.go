package xSolace

type Message struct {
	Subscription string
	Properties   map[string]interface{}
}

type Service interface {
	Publish(msg string, m Message) error
	Listen(fnc func(any)) error
}
