package agent

import (
	"context"
	"fmt"
	sendpb "github.com/cuikai2021/icb-message-agent/agent/go/sendpb"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

type Sender interface {
	SendMessages(messages []*sendpb.Message) error
	Close() error
}

type GRRCSend struct {
	cli   sendpb.TaskMessageServiceClient
	conn  *grpc.ClientConn
	cChan chan struct{} //保证串行调用，确保messages顺序性
}

func NewGRPCSender(serverAddr string) (*GRRCSend, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	creds := credentials.NewClientTLSFromCert(nil, "")
	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithInsecure(), grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	gp := &GRRCSend{
		cli:   sendpb.NewTaskMessageServiceClient(conn),
		conn:  conn,
		cChan: make(chan struct{}, 1),
	}

	return gp, nil
}

func (gp *GRRCSend) SendMessages(messages []*sendpb.Message) error {
	gp.cChan <- struct{}{}
	defer func() {
		<-gp.cChan
	}()

	header := metadata.Pairs("authorization", fmt.Sprintf("bearer %s", os.Getenv("ICB_USER_TOKEN")))
	headerCtx := metadata.NewOutgoingContext(context.Background(), header)
	ctx, _ := context.WithTimeout(headerCtx, 5*time.Second)

	if err := gp.send(ctx, messages); err != nil {
		return err
	}

	return nil
}

func (gp *GRRCSend) send(ctx context.Context, messages []*sendpb.Message) error {
	_, err := gp.cli.SendMessage(ctx, &sendpb.SendMessageRequest{
		TaskId:   os.Getenv("ICB_RUN_ID"), //从环境变量中获取RunId
		Messages: messages,
	})
	if err != nil {
		// 认为服务不可用
		log.Printf("grpc send %s\n", err)
	}

	return err
}

func (gp *GRRCSend) Close() error {
	if gp.conn != nil {
		return gp.conn.Close()
	}
	return nil
}

type LocalSend struct {
	cChan chan struct{} //保证串行调用，确保messages顺序性
}

func NewLocalSender() *LocalSend {
	return &LocalSend{
		cChan: make(chan struct{}, 1),
	}
}

func (l *LocalSend) SendMessages(messages []*sendpb.Message) error {
	l.cChan <- struct{}{}
	defer func() {
		<-l.cChan
	}()

	if err := l.send(messages); err != nil {
		return err
	}

	return nil
}

func (l *LocalSend) send(messages []*sendpb.Message) error {
	for _, message := range messages {
		fmt.Println(fmt.Sprintf("[%s] %s", message.Level.String(), message.Message))
	}

	return nil
}

func (l *LocalSend) Close() error {
	return nil
}
