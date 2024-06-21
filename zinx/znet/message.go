package znet

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

func (msg *Message) GetMsgId() uint32 {
	return uint32(msg.Id)
}
func (msg *Message) GetMsgData() []byte {
	return msg.Data
}
func (msg *Message) GetMsgLen() uint32 {
	return uint32(msg.DataLen)
}
func (msg *Message) SetMsgId(id uint32) {
	msg.Id = id
}
func (msg *Message) SetMsgData(data []byte) {
	msg.Data = data
}
func (msg *Message) SetMsgLen(flag uint32) {
	msg.DataLen = flag
}
func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}
