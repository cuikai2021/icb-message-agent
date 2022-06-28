package agent

import (
	sendpb "github.com/ICBench/icb-message-agent/agent/go/sendpb"
	"log"
	"os"
	"sync"
	"time"
)

type batchSender struct {
	mutex sync.Mutex

	senderType string

	sender Sender
	packet *packet

	once sync.Once
}

func NewBatchSender(senderType string) *batchSender {
	s := &batchSender{
		senderType: senderType,
	}

	s.once.Do(func() {
		go s.initSender()
		go s.backgroundSendPacket()
	})
	return s
}

func (s *batchSender) Send(message *sendpb.Message) (err error) {
	s.mutex.Lock()
	if s.packet == nil {
		s.packet = newPacket()
	}
	p := s.packet.append(message)
	if p.isFull {
		s.send(p)
	}
	s.mutex.Unlock()

	return nil
}

// push 异常情况，比如grpc连接未成功创建或者发送失败，都直接丢弃
func (s *batchSender) send(p *packet) {
	defer func() {
		s.packet = nil
	}()
	if len(p.messages) <= 0 {
		// 没有messages的包，直接丢掉
		return
	}
	if s.sender == nil {
		p.free()
		return
	}

	go func() {
		if err := s.sender.SendMessages(p.messages); err != nil {
			log.Printf("send remote message packet %s\n", err.Error())
		}
		p.free()
	}()
}

func (s *batchSender) initSender() {
	tick := time.NewTicker(time.Second)
	for {
		if s.sender == nil {
			if s.senderType == "local" {
				s.sender = NewLocalSender()
				log.Print("init local sender success \n")
				tick.Stop()
				return
			} else {
				env := os.Getenv("DEPLOY_MODEL")
				serverAddr := "ginkgo.grpc.icbench.com:1443"
				if env == "staging" {
					serverAddr = "ginkgo.internal.icbench.com:1443"
				}

				sender, err := NewGRPCSender(serverAddr)
				if err != nil {
					log.Printf("init grpc push error %s\n", err.Error())
					goto next
				}
				s.sender = sender

				log.Print("init grpc sender success \n")
				tick.Stop()
				return
			}
		}
	next:
		select {
		case <-tick.C:
		}
	}
}

// 当一定时间内，包容量没有达到，则也会默认发送已在缓存中的messages
func (s *batchSender) backgroundSendPacket() {
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			s.mutex.Lock()
			if s.packet != nil {
				p := s.packet.flush()
				s.send(p)
			}
			s.mutex.Unlock()
		}
	}
}
