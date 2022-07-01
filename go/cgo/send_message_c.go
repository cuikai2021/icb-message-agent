package main

import "C"

import (
	agent "github.com/cuikai2021/icb-message-agent"
	proto "github.com/cuikai2021/icb-message-agent/proto"
)

type MessageLevel int32

const (
	MessageLevel_INFO  MessageLevel = 0
	MessageLevel_WARN  MessageLevel = 1
	MessageLevel_ERROR MessageLevel = 2
)

//export SendMessageNew
func SendMessageNew(level MessageLevel, msgTemplate string, message string) {
	agent.SendMessageWithTemplate(proto.MessageLevel(level), msgTemplate, message)
}

func main() {
}
