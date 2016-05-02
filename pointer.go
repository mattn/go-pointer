package pointer

import (
	"unsafe"
)

var (
	index uintptr = 0
	store         = map[uintptr]interface{}{}
)

func Save(v interface{}) unsafe.Pointer {
	store[index] = v
	curr := index
	index++
	return unsafe.Pointer(curr)
}

func Restore(i uintptr) interface{} {
	if v, ok := store[i]; ok {
		return v
	}
	panic("invalid pointer")
}

func Unref(i uintptr) {
	delete(store, i)
}
