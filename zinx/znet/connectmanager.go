package znet

import (
	"errors"
	"fmt"
	"sync"
	"zinx/zinx/ziface"
)

type ConnectManager struct {
	connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

func NewConnectManager() *ConnectManager {
	return &ConnectManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

func (connManager *ConnectManager) Add(conn ziface.IConnection) {
	connManager.connLock.Lock()
	defer connManager.connLock.Unlock()

	connManager.connections[conn.GetConnectID()] = conn
}
func (connManager *ConnectManager) Remove(conn ziface.IConnection) {
	connManager.connLock.Lock()
	defer connManager.connLock.Unlock()

	delete(connManager.connections, conn.GetConnectID())
}
func (connManager *ConnectManager) Get(connID uint32) (ziface.IConnection, error) {
	connManager.connLock.RLock()
	defer connManager.connLock.RUnlock()
	if conn, ok := connManager.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}
func (connManager *ConnectManager) Count() int {
	return len(connManager.connections)
}
func (connManager *ConnectManager) Clear() {
	connManager.connLock.Lock()
	defer connManager.connLock.Unlock()

	for connID, conn := range connManager.connections {
		conn.Stop()
		delete(connManager.connections, connID)
	}
	fmt.Println("clear ok")
}
