package pointer

import (
	"sync"
	"unsafe"
)

var (
	index uintptr = 100
	store         = map[uintptr]interface{}{}
	mutex sync.Mutex
)

func Save(v interface{}) unsafe.Pointer {
	mutex.Lock()
	defer mutex.Unlock()
	store[index] = v
	curr := index
	index++
	return unsafe.Pointer(curr)
}

func Restore(i uintptr) interface{} {
	mutex.Lock()
	defer mutex.Unlock()
	if v, ok := store[i]; ok {
		return v
	}
	panic("invalid pointer")
}

func Unref(i uintptr) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(store, i)
}
