package main

import (
	"context"
	"sync"

	"github.com/nats-io/nats.go"
)

type (
	DummyConsumer struct {
		ctx    context.Context
		waiter *sync.WaitGroup
		nc     *nats.Conn
	}
)

func MakeDummyConsumer(ctx context.Context, nc *nats.Conn, waiter *sync.WaitGroup) DummyConsumer {
	return DummyConsumer{
		ctx:    ctx,
		nc:     nc,
		waiter: waiter,
	}
}

func (bc *DummyConsumer) Run() {
}
