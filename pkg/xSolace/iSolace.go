package xSolace

import "context"

type Pipe struct {
	topic      map[string]string
	MsgOut     chan string
	ReceiptOut chan any
	Err        chan error
	Properties map[string]interface{}
}

type Service interface {
	Publish(ctx context.Context, m *Pipe)
	Listen(ctx context.Context, m *Pipe)
}
