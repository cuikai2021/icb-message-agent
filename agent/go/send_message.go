package agent

import (
	"errors"
	"fmt"
	sendpb "github.com/ICBench/icb-message-agent/agent/go/sendpb"
	"os"
)

var (
	sender  *batchSender
	checker *legalChecker
)

func init() {
	sender = NewBatchSender(os.Getenv("SENDER_TYPE"))

	checker = NewLegalChecker()
}

// SendMessage interface for sendMessage
func SendMessage(level sendpb.MessageLevel, format string, args ...interface{}) (err error) {

	if !checker.IsLegal(format) {
		return errors.New(fmt.Sprintf("msg %s should be reviewed before send", format))
	}

	return sender.Send(&sendpb.Message{
		Level:   level,
		Message: fmt.Sprintf(format, args...),
	})
}
