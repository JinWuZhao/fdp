package simpleudp

import (
	"errors"
	"sync/atomic"
)

type pktBuffer struct {
	ring []packet
	head int32
	tail int32
	cap  int32
}

func newPktBuffer(cap int32) *pktBuffer {
	return &pktBuffer{
		ring: make([]packet, cap),
		head: cap - 1,
		tail: cap - 1,
		cap:  cap,
	}
}

func (b *pktBuffer) Push(pkt packet) error {
	head := (atomic.AddInt32(&b.head, -1) + 1) % b.cap
	if head < 0 {
		head += b.cap
	}
	if b.ring[head] != nil {
		return errors.New("buffer is almost full")
	}
	b.ring[head] = pkt
	return nil
}

func (b *pktBuffer) Pop() (packet, error) {
	tail := (atomic.AddInt32(&b.tail, -1) + 1) % b.cap
	if tail < 0 {
		tail += b.cap
	}
	pkt := b.ring[tail]
	if pkt == nil {
		return nil, errors.New("buffer is almost empty")
	}
	b.ring[tail] = nil
	return pkt, nil
}

func (b *pktBuffer) Get(index int32) (packet, error) {
	head := atomic.LoadInt32(&b.head) % b.cap
	if head < 0 {
		head += b.cap
	}
	ringIndex := head + index
	if ringIndex >= b.cap {
		ringIndex -= b.cap
	}
	pkt := b.ring[ringIndex]
	if pkt == nil {
		return nil, errors.New("index out of range")
	}
	return pkt, nil
}
