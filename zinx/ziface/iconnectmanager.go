package ziface

// IConnectManager 管理conn(s)
type IConnectManager interface {
	Add(conn IConnection)
	Remove(conn IConnection)
	Get(connID uint32) (IConnection, error)
	Count() int
	Clear()
}
