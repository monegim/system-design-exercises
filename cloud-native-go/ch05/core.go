package main

import (
	"errors"
	"sync"
)

// var store = make(map[string]string)

var store = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

var ErrorNoSuchKey = errors.New("no such key")

func Get(key string) (string, error) {
	value, ok := store[key]
	if !ok {
		return "", ErrorNoSuchKey
	}
	return value, nil
}

func Put(key string, value string) error {
	store[key] = value
	return nil
}

func Delete(key string) error {
	delete(store, key)
	return nil
}
