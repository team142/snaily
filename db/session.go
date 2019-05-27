package db

import (
	"sync"
	"time"
)

var GlobalSessionCache = buildSessionCache()

func buildSessionCache() *SessionCache {
	result := &SessionCache{
		store: make(map[string]SessionRow),
	}
	return result
}

type SessionCache struct {
	store map[string]SessionRow
	m     sync.Mutex
}

type SessionRow struct {
	Key        string
	UserID     string
	ValidUntil time.Time
}

func (s *SessionCache) SessionValid(key string) bool {
	s.m.Lock()
	defer s.m.Unlock()
	if r, ok := s.store[key]; ok {
		return r.ValidUntil.After(time.Now())
	}
	return false

}

func (s *SessionCache) SetSession(key, ID string, duration time.Duration) {
	s.m.Lock()
	defer s.m.Unlock()
	r := SessionRow{
		Key:        key,
		UserID:     ID,
		ValidUntil: time.Now().Add(duration),
	}
	s.store[key] = r
	return
}
