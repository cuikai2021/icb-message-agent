package agent

import (
	proto "github.com/cuikai2021/icb-message-agent/proto"
	"sync"
)

var (
	_maxPacketSize = 100

	_packetPool = sync.Pool{New: func() interface{} {
		return &packet{
			messages: make([]*proto.Message, 0),
		}
	}}
)

type packet struct {
	messages []*proto.Message
	isFull   bool
}

func (p *packet) free() {
	p.messages = p.messages[:0]
	p.isFull = false
	_packetPool.Put(p)
}

func (p *packet) append(msg *proto.Message) *packet {
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
