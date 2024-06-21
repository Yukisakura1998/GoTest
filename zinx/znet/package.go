package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/zinx/utils"
	"zinx/zinx/ziface"
)

type Package struct {
	DataLen uint32
	Id      uint32
	Data    []byte
}

func (pkg *Package) Pack(msg ziface.IMessage) ([]byte, error) {
	buff := bytes.NewBuffer([]byte{})
	if err := binary.Write(buff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(buff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	if err := binary.Write(buff, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}
func (pkg *Package) Unpack(data []byte) (ziface.IMessage, error) {
	readResult := bytes.NewReader(data)
	msg := &Message{}
	if err := binary.Read(readResult, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(readResult, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("package too big")
	}
	return msg, nil
}
func (pkg *Package) GetHeadLen() uint32 {
	//head + id ,4 + 4
	return 8
}

func NewPackage() *Package {
	return &Package{}
}
