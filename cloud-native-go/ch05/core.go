package main

var store = make(map[string]string)

func Put(key string, value string) error {
	store[key] = value
	return nil
}
