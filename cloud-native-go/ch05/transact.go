package main

import (
	"fmt"
	"os"
	"sync"
)

type EventType byte

const (
	_                     = iota
	EventDelete EventType = iota
	EventPut
)

type Event struct {
	Sequence  uint64
	EventType EventType
	Key       string
	Value     string
}

type TransactionLogger struct {
	events       chan<- Event
	errors       <-chan error
	lastSequence uint64
	file         *os.File
	wg           *sync.WaitGroup
}

func (l *TransactionLogger) WritePut(key, value string) {
	l.wg.Add(1)
	l.events <- Event{EventType: EventDelete, Key: key}
}

func (l *TransactionLogger) WriteDelete(key string) {
	l.wg.Add(1)
	l.events <- Event{EventType: EventDelete, Key: key}
}

func (l *TransactionLogger) Err() <-chan error {
	return l.errors
}

func NewTransactionLogger(filename string) (*TransactionLogger, error) {
	var err error
	var l TransactionLogger = TransactionLogger{wg: &sync.WaitGroup{}}
	l.file, err = os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return nil, fmt.Errorf("cannot open transaction log file: %w", err)
	}
	return &l, nil
}

func (l *TransactionLogger) Run() {
	events := make(chan Event, 16)
	l.events = events

	errors := make(chan error, 1)
	l.errors = errors

	go func() {
		for e := range events {
			l.lastSequence++
			_, err := fmt.Fprintf(
				l.file,
				"%d\t%d\t%s\t%s\n",
				l.lastSequence, e.EventType, e.Key, e.Value)
			if err != nil {
				errors <- fmt.Errorf("cannot write to log file: %w", err)
			}
			l.wg.Done()
		}
	}()
}

func (l *TransactionLogger) Wait() {
	l.wg.Wait()
}

func (l *TransactionLogger) Close() error {
	l.wg.Wait()

	if l.events != nil {
		close(l.events)
	}
	return l.file.Close()
}
