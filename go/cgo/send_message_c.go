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

//export SendMessage
func SendMessage(level MessageLevel, msgTemplate string, message string) {
	msgGo := make([]byte, len(message))
	copy(msgGo, message)
	msgTemplateGo := make([]byte, len(msgTemplate))
	copy(msgTemplateGo, msgTemplateGo)
	agent.SendMessageWithTemplate(proto.MessageLevel(level), string(msgTemplateGo), string(msgGo))
}

func main() {
}
