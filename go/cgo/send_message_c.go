package main

import "C"

import (
	agent "github.com/cuikai2021/icb-message-agent"
	proto "github.com/cuikai2021/icb-message-agent/proto"
	"time"
)

//export SendMessage
func SendMessage() {
	format := "This is a test legal message %s %s"
	for i := 0; i < 10; i++ {
		agent.SendMessageWithTemplate(proto.MessageLevel_INFO, format, format)
	}
	time.Sleep(time.Second * 5)
}

//import (
//	agent "github.com/cuikai2021/icb-message-agent"
//	proto "github.com/cuikai2021/icb-message-agent/proto"
//)

type MessageLevel int32

const (
	MessageLevel_INFO  MessageLevel = 0
	MessageLevel_WARN  MessageLevel = 1
	MessageLevel_ERROR MessageLevel = 2
)

//export SendMessageNew
func SendMessageNew(level MessageLevel, msgTemplate string, message string) {
	for i := 0; i < 10; i++ {
		agent.SendMessageWithTemplate(proto.MessageLevel(level), msgTemplate, message)
	}
	time.Sleep(time.Second * 5)
}

func main() {
}
