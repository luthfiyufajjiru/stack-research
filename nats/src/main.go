package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	AppName = "explore-nats"
)

func setLog() {
	logPrefix := fmt.Sprintf("[%s] ", AppName)
	log.SetPrefix(logPrefix)
}

func quitSignal() chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	return quit
}

func exit(cfg *Configuration, waiter *sync.WaitGroup) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ClosingTimeout)
	defer cancel()

	runnerSignal := make(chan any)
	go func() {
		waiter.Wait()
		close(runnerSignal)
	}()

	select {
	case <-ctx.Done():
		log.Println("shutdown timeout reached, exiting")
	case <-runnerSignal:
		log.Println("all runners finished, exiting")
	}
}

func main() {
	setLog()

	cfg := new(Configuration)
	if err := cfg.Load(); err != nil {
		log.Panic("failed to load configuration")
	}

	ctx, cancel := context.WithCancel(context.Background())

	nc, err := EstablishNatsConnection(cfg)
	if err != nil {
		log.Panic("failed to connect to nats server")
	}

	var runners sync.WaitGroup

	quit := quitSignal()

	for _, runner := range cfg.Runners {
		switch runner {
		case RunnerDummyConsumer:
			runners.Add(1)
			dummyConsumer := MakeDummyConsumer(ctx, nc, &runners)
			go dummyConsumer.Run()
		case RunnerDummyPublisher:
			runners.Add(1)
			dummyPublisher := MakeDummyPublisher(ctx, nc, &runners)
			go dummyPublisher.Run()
		}
	}

	<-quit
	cancel()

	defer nc.Close()
	defer nc.Drain()

	exit(cfg, &runners)
}
