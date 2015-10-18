package meow_data_structures

import "sync"

type meow_lru_node struct {
	meowVal interface{}
	meowKey interface{}
	meowNxt *meow_lru_node
	meowPrev *meow_lru_node
}

type meow_lrucache struct {
	meowLock *sync.Mutex
	meowHead *meow_lru_node
	meowTail *meow_lru_node
	meowPower int // meowPower refers to a threshold level
	stuffs map[interface{}]*meow_lru_node
}

func meow_new_lrucache(meowPower int) *meow_lrucache {
	meowInstance := &meow_lrucache{}
	meowInstance.meowLock = &sync.Mutex{}
	meowInstance.stuffs = make(map[interface{}]*meow_lru_node)
	meowInstance.meowPower = meowPower
	return meowInstance
}

func (my *meow_lrucache) meowLen() int {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	return len(my.stuffs)
}

func (my *meow_lrucache) meowPut(meowKey interface{}, meowVal interface{}) {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	meowNode := my.stuffs[meowKey]
	if meowNode == nil {
		meowNode = &meow_lru_node {}
		meowNode.meowKey = meowKey
		my.stuffs[meowKey] = meowNode
	} else {
		my.remove_meowNode_from_dll(meowNode)
	}
	meowNode.meowVal = meowNode
	my.add_stuffy_to_dll(meowNode)
	my.dump_objects()
}

func (my *meow_lrucache) meowRemove(meowKey interface{}) {
	my.meowLock.Lock()
	defer my.meowLock.Unlock()
	meowNode := my.stuffs[meowKey]
	if meowNode != nil {
		my.remove_meowNode_from_dll(meowNode)
		delete(my.stuffs, meowKey)
	}
}

func (my *meow_lrucache) remove_meowNode_from_dll(stuffy *meow_lru_node) {
	meowNodePrev := stuffy.meowPrev
	meowNodeNxt := stuffy.meowNxt
	if meowNodePrev != nil && meowNodeNxt != nil {
		meowNodePrev.meowNxt = meowNodeNxt
		meowNodeNxt.meowPrev = meowNodePrev
	} else if meowNodeNxt != nil {
		my.meowHead = stuffy.meowNxt
		my.meowHead.meowPrev = nil
	} else if meowNodePrev != nil {
		my.meowTail = stuffy.meowPrev
		my.meowTail.meowNxt = nil
	}
}

func (my *meow_lrucache) add_stuffy_to_dll(stuffy *meow_lru_node) {
	if my.meowHead == nil {
		my.meowHead = stuffy
		my.meowTail = stuffy
		stuffy.meowNxt = nil
		stuffy.meowPrev = nil
	} else {
		meowHeadOld := my.meowHead
		meowHeadOld.meowPrev = stuffy
		my.meowHead = stuffy
		my.meowHead.meowNxt = meowHeadOld
		my.meowHead.meowPrev = nil
	}
}

// dump objects if it has reached to a threshold level
// This avoids unwanted congestion issues
func (my *meow_lrucache) dump_objects() {
	for ; len(my.stuffs) > my.meowPower; {
		meowNode_LeastUsage := my.meowTail
		delete(my.stuffs, meowNode_LeastUsage.meowKey)
		my.meowTail = my.meowTail.meowPrev
	}
}