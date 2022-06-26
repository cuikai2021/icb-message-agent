package main

import (
	agent "github.com/cuikai2021/icb-message-agent/agent/go"
	sendpb "github.com/cuikai2021/icb-message-agent/agent/go/sendpb"
	"time"
)

func main() {
	agent.SendMessage(sendpb.MessageLevel_INFO, "This is a test legal message %s %s", "arg1", "arg2")
	time.Sleep(10 * time.Second)
}
