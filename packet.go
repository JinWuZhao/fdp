package fdp

import (
	"fmt"
	"strings"
)

type packet interface {
	Marshal() []byte
	Unmarshal([]byte) error
	fmt.Stringer
}

const (
	idEnd               = 16
	lastIDEnd           = idEnd + 16
	dataLenEnd          = lastIDEnd + 16
	dataPacketHeaderLen = dataLenEnd
)

type dataPacket struct {
	ID      []byte
	LastID  []byte
	DataLen uint16
	Data    []byte
}

func newDataPacket() *dataPacket {
	return &dataPacket{}
}

func (p *dataPacket) Marshal() []byte {
	return nil
}

func (p *dataPacket) Unmarshal([]byte) error {
	return nil
}

func (p *dataPacket) String() string {
	return fmt.Sprintf("dataPacket: ID = %s, LastID = %s, DataLen = %d, Data: %v", string(p.ID), string(p.LastID), p.DataLen, p.Data)
}

const (
	numEnd = 16
	idLen  = 16
)

type ackPacket struct {
	Num uint16
	IDS [][]byte
}

func newAckPacket() *ackPacket {
	return &ackPacket{}
}

func (p *ackPacket) Marshal() []byte {
	return nil
}

func (p *ackPacket) Unmarshal([]byte) error {
	return nil
}

func (p *ackPacket) String() string {
	ids := make([]string, len(p.IDS))
	for i, v := range p.IDS {
		ids[i] = string(v)
	}
	return fmt.Sprintf("ackPacket: Num = %d, IDS = [ %s ]", p.Num, strings.Join(ids, " "))
}
