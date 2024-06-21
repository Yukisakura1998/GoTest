package ziface

type IPackage interface {
	Pack(msg IMessage) ([]byte, error)
	Unpack([]byte) (IMessage, error)
	GetHeadLen() uint32
}
