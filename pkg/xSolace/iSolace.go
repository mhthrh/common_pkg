package xSolace

import "context"

type Pipe struct {
	Topic      string
	MsgOut     chan string
	MsgIn      chan string
	ReceiptOut chan any
	Err        chan error
	Properties map[string]interface{}
}

type Service interface {
	Publish(ctx context.Context, m *Pipe)
	Listen(ctx context.Context, m *Pipe)
}
