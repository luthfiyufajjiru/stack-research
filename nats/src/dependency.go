package main

import "github.com/nats-io/nats.go"

func EstablishNatsConnection(cfg *Configuration) (nc *nats.Conn, err error) {
	nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		return
	}
	return
}
