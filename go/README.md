# go version

## Import

```go
  import (
	agent "github.com/ICBench/icb-message-agent/agent/go"
	sendpb "github.com/ICBench/icb-message-agent/agent/go/sendpb"
  )
```

## Interface

Same as `fmt.Printf`

```go
  agent.SendMessage(sendpb.MessageLevel_INFO, "This is a test legal message %s %s", "arg1", "arg2")
```
