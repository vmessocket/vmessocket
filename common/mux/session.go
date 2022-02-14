package mux

import (
	"sync"

	"github.com/vmessocket/vmessocket/common"
	"github.com/vmessocket/vmessocket/common/buf"
	"github.com/vmessocket/vmessocket/common/protocol"
)

type Session struct {
	input        buf.Reader
	output       buf.Writer
	parent       *SessionManager
	ID           uint16
	transferType protocol.TransferType
}

type SessionManager struct {
	sync.RWMutex
	sessions map[uint16]*Session
	count    uint16
	closed   bool
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		count:    0,
		sessions: make(map[uint16]*Session, 16),
	}
}

func (m *SessionManager) Add(s *Session) {
	m.Lock()
	defer m.Unlock()
	if m.closed {
		return
	}
	m.count++
	m.sessions[s.ID] = s
}

func (s *Session) Close() error {
	common.Close(s.output)
	common.Close(s.input)
	return nil
}

func (m *SessionManager) Close() error {
	m.Lock()
	defer m.Unlock()
	if m.closed {
		return nil
	}
	m.closed = true
	for _, s := range m.sessions {
		common.Close(s.input)
		common.Close(s.output)
	}
	m.sessions = nil
	return nil
}

func (m *SessionManager) Closed() bool {
	m.RLock()
	defer m.RUnlock()
	return m.closed
}

func (m *SessionManager) Count() int {
	m.RLock()
	defer m.RUnlock()
	return int(m.count)
}

func (m *SessionManager) Get(id uint16) (*Session, bool) {
	m.RLock()
	defer m.RUnlock()
	if m.closed {
		return nil, false
	}
	s, found := m.sessions[id]
	return s, found
}
