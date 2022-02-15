package mux

import (
	"sync"

	"github.com/vmessocket/vmessocket/common"
	"github.com/vmessocket/vmessocket/common/buf"
)

type Session struct {
	input        buf.Reader
	output       buf.Writer
	ID           uint16
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
