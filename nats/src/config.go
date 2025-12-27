package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const (
	RunnerDummyConsumer  = "dummy_consumer"
	RunnerDummyPublisher = "dummy_publisher"
)

type (
	Configuration struct {
		Runners        []string
		ClosingTimeout time.Duration
		Nats           NatsConfig
	}

	NatsConfig struct {
	}
)

func (c *Configuration) Load() (err error) {
	if err = godotenv.Load(); err != nil {
		log.Println("env file not found: using os host env")
		err = nil
	}

	if runnersStr := os.Getenv("RUNNERS"); runnersStr != "" {
		c.Runners = strings.Split(runnersStr, ",")
	}

	if len(c.Runners) == 0 {
		c.Runners = []string{RunnerDummyConsumer, RunnerDummyPublisher}
	}

	return
}
