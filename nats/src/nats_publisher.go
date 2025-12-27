package main

import (
	"context"
	"sync"

	"github.com/nats-io/nats.go"
)

type (
	DummyPublisher struct {
		ctx    context.Context
		nc     *nats.Conn
		waiter *sync.WaitGroup
	}
)

func MakeDummyPublisher(ctx context.Context, nc *nats.Conn, waiter *sync.WaitGroup) DummyPublisher {
	return DummyPublisher{
		ctx:    ctx,
		nc:     nc,
		waiter: waiter,
	}
}

func (bp *DummyPublisher) Run() {}
