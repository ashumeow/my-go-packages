// package
package meow_data_structures

// importing other packages
import (
	"fmt"
	"math/rand"
	"sync"
	"bytes"
	"time"
)

// type
type meowArrayList struct {
	meowCount int
	meowLock *sync.Mutex
	stuffs []interface{}
}

// meowNewArrayList
func meowNewArrayList() *meowArrayList {
	meowInstance := &meowArrayList {}
	meowInstance.meowLock = &sync.Mutex{}
	meowInstance.stuffs = make([]interface{}, 10)
	meowInstance.meowCount = 0
	rand.Seed(time.Now().UTC().UnixNano())
	return meowInstance
}

// meowLen
func (my *meowArrayList) meowLen() int {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return my.meowCount
}

// if empty
func (my *meowArrayList) meowEmpty() bool {
	return my.meowLen() == 0
}

// add
func (my *meowArrayList) meowAdd(objects ...interface{}) {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()

	for o := range objects {
		my.meowAdd(o)
	}
}
func (my *meowArrayList) meow_add(o interface{}) {
	my.stuffs[my.meowCount] = o
	my.meowCount++
	my.resize()
}

// resize if required
func (my *meowArrayList) resize() {
	// adding capacity
	meowPower := cap(my.stuffs)

	if(my.meowCount >= (meowPower - 1)) {
		// init new capacity
		meowPowerUp := (meowPower + 1) * 2
		// init temp
		temp := make([]interface{}, meowPowerUp, meowPowerUp)
		copy(temp, my.stuffs)
	}
}

// slicing
func (my *meowArrayList) meowSlice() []interface{} {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	out := make([]interface{}, my.meowCount)
	copy(out, my.stuffs)
	return out
}

// indexing
func (my *meowArrayList) meowIndex(o interface{}) int {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return meow_index(o)
}
func (my *meowArrayList) meow_index(o interface{}) int {
	meow_indexer := -1
	for x := 0; x < my.meowCount; x++ {
		if my.stuffs[x] == 0 {
			meow_indexer = x
			break;
		}
	}
	return meow_indexer
}

// fetching
func (my *meowArrayList) meowFetch(meow_indexer int) interface{} {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return my.stuffs[meow_indexer]
}

// sampling
func (my *meowArrayList) meowSample() interface{} {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	if(my.meowCount == 0) {
		return nil
	}
	meow_indexer := rand.Intn(my.meowCount)
	return my.stuffs[meow_indexer]
}

// registers
func (my *meowArrayList) meowRegisters(o interface{}) bool {
	return my.meowIndex(o) != -1
}

// remove
func (my *meowArrayList) meowRemove(o interface{}) bool {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	meow_indexer := my.meow_index(o)
	if meow_indexer == -1 {
		return false
	}
	my.stuffs[meow_indexer] = nil
	for x := meow_indexer; x < my.meowCount - 1; x++ {
		my.meow_swap(x, x+1)
	}
	my.meowCount--
	return true
}

// swapping
func (my *meowArrayList) meowSwap(y int, yy int) {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	my.meow_swap(y, yy)
}
func (my *meowArrayList) meow_swap(y int, yy int) {
	my.stuffs[y], my.stuffs[yy] = my.stuffs[yy], my.stuffs[y]
}

// reset
func (my *meowArrayList) meowReset() {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	meowPower := cap(my.stuffs)
	meowLength := len(my.stuffs)
	my.stuffs = make([]interface{}, meowLength, meowPower)
	my.meowCount = 0
}

// adding from meowArrayList
func (my *meowArrayList) meowAddFromList(meowList *meowArrayList) {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	if meowList == nil {
		return
	}
	for x := 0; x < meowList.meowLen(); x++ {
		my.meow_add(meowList.meowFetch(x))
	}
}

// oh yeah... source..
func (my *meowArrayList) meowSource() interface{} {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return my.stuffs[0]
}

// destination
func (my *meowArrayList) meowDest() interface{} {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return my.stuffs[my.meowCount - 1]
}

// strings
func (my *meowArrayList) meowString() string {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	var meowBuffer bytes.Buffer
	for x := 0; x < my.meowCount; x++ {
		stuffy := my.stuffs[x]
		meowStrfy := fmt.Sprintf("%s", stuffy)
		meowBuffer.WriteString(meowStrfy)
		if x != (my.meowCount-1) {
			meowBuffer.WriteString(", ")
		}
	}
	return fmt.Sprintf("[ %s ]", meowBuffer.meowString())
}