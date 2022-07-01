package agent

import (
	"errors"
	"fmt"
	proto "github.com/cuikai2021/icb-message-agent/proto"
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
func SendMessage(level proto.MessageLevel, format string, args ...interface{}) (err error) {

	if !checker.IsLegal(format) {
		return errors.New(fmt.Sprintf("msg %s should be reviewed before send", format))
	}

	return sender.Send(&proto.Message{
		Level:   level,
		Message: fmt.Sprintf(format, args...),
	})
}

func SendMessageWithTemplate(level proto.MessageLevel, format string, msg string) (err error) {

	if !checker.IsLegal(format) {
		return errors.New(fmt.Sprintf("msg %s should be reviewed before send", format))
	}

	return sender.Send(&proto.Message{
		Level:   level,
		Message: msg,
	})
}
