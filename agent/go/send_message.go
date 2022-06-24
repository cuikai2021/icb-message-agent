package client

import (
	"errors"
	"fmt"
	sendpb "github.com/cuikai2021/icb-message-agent/agent/go/sendpb"
)

var (
	sender  *batchSender
	checker *legalChecker
)

func init() {
	//env := os.Getenv("DEPLOY_MODEL")
	//if env == "staging" {
	//	sender = NewBatchSender("ginkgo.internal.icbench.com:1443")
	//} else {
	//	sender = NewBatchSender("ginkgo.grpc.icbench.com:1443")
	//}

	sender = NewBatchSender("127.0.0.1:8776")

	checker = NewLegalChecker()
}

// SendMessage interface for sendMessage
func SendMessage(level sendpb.MessageLevel, format string, args ...interface{}) (err error) {

	if !checker.IsLegal(format) {
		return errors.New(fmt.Sprintf("msg %s should be reviewed before send", format))
	}

	msg := fmt.Sprintf(format, args)
	return sender.Send(&sendpb.Message{
		Level:   level,
		Message: msg,
	})
}
