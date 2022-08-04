package main

import (
	"log"
	"os"
	"time"

	"github.com/sony/gobreaker"
)

type ClientCircuitBreakerProxy struct {
	client NotificationClient
	logger *log.Logger
	gb     *gobreaker.CircuitBreaker
}

func shouldBeSwitchedToOpen(counts gobreaker.Counts) bool {
	failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
	return counts.Requests >= 3 && failureRatio >= 0.6
}

func NewClientBreakerProxy(client NotificationClient) *ClientCircuitBreakerProxy {
	logger := log.New(os.Stdout, "CB\t", log.LstdFlags)

	cfg := gobreaker.Settings{
		Interval:    5 * time.Second,
		Timeout:     7 * time.Second,
		ReadyToTrip: shouldBeSwitchedToOpen,
		OnStateChange: func(_ string, from gobreaker.State, to gobreaker.State) {
			logger.Println("state changed from", from.String(), "to", to.String())
		},
	}
	return &ClientCircuitBreakerProxy{
		client: client,
		logger: logger,
		gb:     gobreaker.NewCircuitBreaker(cfg),
	}
}

func (c *ClientCircuitBreakerProxy) Send() error {
	_, err := c.gb.Execute(func() (interface{}, error) {
		err := c.client.Send()
		return nil, err
	})
	return err
}
