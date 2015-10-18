// meow_queue.go

package meow-data-structures

import "sync"

type meowQueueNode struct {
	meowData interface{}
	meowNxt *meowQueueNode
}

// go-routine safe FIFO
type meowQueue struct {
	meowHead *meowQueueNode
	meowTail *meowQueueNode
	meowCount int
	meowLock *sync.Mutex
}

// creating a new pointer to a new meowQueue
func meowNewQueue() *meowQueue {
	q := &meowQueue{}
	q.meowLock = &sync.Mutex{}
	return q
}

// go-routine safe...
func (q *meowQueue) meowLen() int {
	q.meowLock.Lock()
	defer q.meowLock.Unlock()
	// returns number of elements in the meowQueue
	// ...in terms of size/length
	return q.meowCount
}

// meowQueue mutation
// go-routine safe
// meowPush <==> more like insertion thing
func (q *meowQueue) meowPush(stuffy interface{}) {
	q.meowLock.Lock()
	defer q.meowLock.Unlock()
	// pushing/inserting a value at the meowTail (i.e. end) of the meowQueue
	ig := &meowQueueNode{meowData: stuffy}
	if q.meowTail == nil {
		q.meowTail = ig
		q.meowHead = ig
	} else {
		q.meowTail.meowNxt = ig
		q.meowTail = ig
	}
	q.meowCount++
}

// FIFO <==> returns
// Queue mutation
func (q *meowQueue) meowDistro() interface() {
	q.meowLock.Lock()
	defer q.meowLock.Unlock()
	if q.meowHead == nil {
		return nil
	}
	ig := q.meowHead
	q.meowHead = ig.meowNxt
	if q.meowHead == nil {
		q.meowTail = nil
	}
	q.meowCount--
	return ig.meowData
}

// No meowQueue mutation
// FIFO <==> reads
func (q *meowQueue) meowErect() interface{} {
	q.meowLock.Lock()
	defer q.meowLock.Unlock()
	ig := q.meowHead
	if ig == nil || ig.meowData == nil {
		return nil
	}
	return ig.meowData
}