package vnet

import "sync"

type IConnectionMgr interface {
	Add(conn IConnection)
	Delete(id int64)
}

type ConnectionMgr struct {
	sync.RWMutex
	connectionMap map[int64]IConnection
}

func NewConnectionMgr() *ConnectionMgr {
	return &ConnectionMgr{
		connectionMap: make(map[int64]IConnection),
	}
}

func (mgr *ConnectionMgr) Add(conn IConnection) {
	mgr.Lock()
	defer mgr.Unlock()

	mgr.connectionMap[conn.GetId()] = conn
}

func (mgr *ConnectionMgr) Delete(id int64) {
	mgr.Lock()
	defer mgr.Unlock()

	delete(mgr.connectionMap, id)
}
