package main

import "testing"

func TestPut(t *testing.T) {
	const key = "create-key"
	const value = "create-value"

	var val interface{}
	var contains bool

	defer delete(store, key)
	_, contains = store[key]
	if contains {
		t.Error("key/value already exists")
	}
	err := Put(key, value)
	if err != nil {
		t.Error(err)
	}
	val, contains = store[key]
	if !contains {
		t.Error("create failed")
	}
	if val != value {
		t.Error("val/value mismatch")
	}

}
