package client

import (
	sendpb "github.com/cuikai2021/icb-message-agent/agent/go/sendpb"
	"sync"
)

var (
	_maxPacketSize = 100

	_packetPool = sync.Pool{New: func() interface{} {
		return &packet{
			messages: make([]*sendpb.Message, _maxPacketSize),
		}
	}}
)

type packet struct {
	messages []*sendpb.Message
	isFull   bool
}

func (p *packet) free() {
	p.messages = p.messages[:0]
	p.isFull = false
	_packetPool.Put(p)
}

func (p *packet) append(msg *sendpb.Message) *packet {
	p.messages = append(p.messages, msg)
	if _maxPacketSize <= len(p.messages) {
		p.isFull = true
	}
	return p
}

func (p *packet) flush() *packet {
	p.isFull = true
	return p
}

func newPacket() *packet {
	p := _packetPool.Get().(*packet)
	p.isFull = false
	return p
}
