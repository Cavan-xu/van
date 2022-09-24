package vnet

import "sync"

type ConnectionMgr struct {
	sync.RWMutex
	connectionMap map[int64]*Connection
}

func NewConnectionMgr() *ConnectionMgr {
	return &ConnectionMgr{
		connectionMap: make(map[int64]*Connection),
	}
}

func (mgr *ConnectionMgr) Add(conn *Connection) {
	mgr.Lock()
	defer mgr.Unlock()

	mgr.connectionMap[conn.id] = conn
}

func (mgr *ConnectionMgr) Delete(id int64) {
	mgr.Lock()
	defer mgr.Unlock()

	delete(mgr.connectionMap, id)
}
