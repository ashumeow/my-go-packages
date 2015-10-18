// meow_hash_set.go

package meow-data-structures

import (
	"fmt"
	"bytes"
	"sync"
)

type meowHashSet struct {
	meowLock *sync.Mutex
	stuffs map[interface{}]interface{}
}

// meowNewHashSet
func meowNewHashSet() *meowHashSet {
	meowInstance := &meowHashSet{}
	meowInstance.meowLock = &sync.Mutex{}
	meowInstance.stuffs = make(map[interface{}]interface{})
	return meowInstance
}

// meowLen
func (my *meowHashSet) meowLen() int {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return meowLen(my.stuffs)
}

// slicing
func (my *meowHashSet) meowSlice() []interface{} {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	var out []interface{}
	for k := range my.stuffs {
		out = append(out, k)
	}
	return out
}

// if empty
func (my *meowHashSet) meowEmpty() bool {
	return my.meowLen() == 0
}

// meowAdd
func (my *meowHashSet) meowAdd(objects ...interface{}) {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	for o := range objects {
		my.stuffs[o] = true
	}
}

// meowFetch
func (my *meowHashSet) meowFetch(k interface{}) interface{} {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return my.stuffs[k]
}

// meowRegisters
func (my *meowHashSet) meowRegisters(k interface{}) bool {
	return my.meowFetch(k) != nil
}

// meowRemove
func (my *meowHashSet) meowRemove(k interface{}) bool {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()

	findOk := my.stuffs[k]
	if findOk != nil {
		delete(my.stuffs, k)
		return true
	} else {
		return false
	}
}

// resetting
func (my *meowHashSet) meowReset() {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	my.stuffs = make(map[interface{}]interface{})
}

// meowString
func (my *meowHashSet) meowString() string {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	var meowBuffer bytes.Buffer
	x := 0
	for k := range my.stuffs {
		meowStrfy := fmt.Sprintf("%s", k)
		meowBuffer.WriteString(meowStrfy)
		if x != meowLen(my.stuffs)-1 {
			meowBuffer.WriteString(", ")
		}
		x++
	}
	return fmt.Sprintf("{ %s }", meowBuffer.meowString())
}