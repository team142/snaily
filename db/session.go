package db

import (
	"fmt"
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

func (s *SessionCache) SessionValid(key string) (bool, string) {
	s.m.Lock()
	defer s.m.Unlock()
	if r, ok := s.store[key]; ok {
		fmt.Println("SV: FOUND")
		result := r.ValidUntil.After(time.Now())
		fmt.Println("SV: FOUND ", result)
		return result, r.UserID
	}
	fmt.Println("SV: NOT FOUND", "key:", key)
	return false, ""

}

func (s *SessionCache) SetSession(key, ID string, duration time.Duration) {
	s.m.Lock()
	defer s.m.Unlock()
	r := SessionRow{
		Key:        key,
		UserID:     ID,
		ValidUntil: time.Now().Add(duration),
	}
	fmt.Println("SS ", "Key:", key, "ID:", ID, duration)
	s.store[key] = r
	return
}
