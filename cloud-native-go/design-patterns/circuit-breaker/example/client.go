package main

import (
	"errors"
	"net/http"
)

type NotificationClient interface {
	Send() error
}

type SmsClient struct {
	baseUrl string
}

func NewSmsClient(baseUrl string) *SmsClient {
	return &SmsClient{
		baseUrl: baseUrl,
	}
}

func (s *SmsClient) Send() error {
	url := s.baseUrl + "/"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errors.New("bad response")
	}
	return nil
}
