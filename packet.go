package simpleudp

type packet interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
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

const (
	numEnd = 16
	idLen  = 16
)

type ackPacket struct {
	Num uint16
	IDS [][]byte
}
