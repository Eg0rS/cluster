package microservices

import "auth/microservices"

func NewFakeLoggerSlack() microservices.LoggerSlackService {
	return &FakeLoggerSlack{}
}

type FakeLoggerSlack struct {
}

func (c *FakeLoggerSlack) Send(sendMessage *microservices.LoggerSlackTemplateDTO) error {
	return nil
}
