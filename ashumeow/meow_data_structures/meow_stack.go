package meow_data_structures

import "sync"

type meowStackNode struct {
	meowData interface{}
	meowNxt *meowStackNode
}

type meowStack struct {
	meowHead *meowStackNode
	meowCount int
	meowLock *sync.Mutex
}

// meowNewStack
func meowNewStack() *meowStack {
	s := &meowStack{}
	s.meowLock = &sync.Mutex{}
	return s
}

// meowLen
func (s *meowStack) meowLen() int {
	s.meowLock.Lock()
	defer s.meowLock.Unlock()
	return s.meowCount
}

// meowPush
func (s *meowStack) meowPush(stuffy interface{}) {
	s.meowLock.Lock()
	defer s.meowLock.Unlock()
	ig := &meowStackNode { meowData: stuffy }
	if s.meowHead == nil {
		s.meowHead = ig
	} else {
		ig.meowNxt = s.meowHead
		s.meowHead = ig
	}
	s.meowCount++
}

// meowJump
func (s *meowStack) meowJump() interface{} {
	s.meowLock.Lock()
	defer s.meowLock.Unlock()
	var ig *meowStackNode
	if s.meowHead != nil {
		ig = s.meowHead
		s.meowHead = ig.meowNxt
		s.meowCount--
	}
	if ig == nil {
		return nil
	}
	return ig.meowData
}

// meowErect
func (s *meowStack) meowErect() interface{} {
	s.meowLock.Lock()
	defer s.meowLock.Unlock()
	ig := s.meowHead
	if ig == nil || ig.meowData == nil {
		return nil
	}
	return ig.meowData
}