package ziface

type IMessage interface {
	GetMsgId() uint32
	GetMsgData() []byte
	GetMsgLen() uint32
	SetMsgId(id uint32)
	SetMsgData(data []byte)
	SetMsgLen(flag uint32)
}
