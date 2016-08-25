package pointer

// #include <stdlib.h>
import "C"
import (
	"sync"
	"unsafe"
)

var (
	store = map[unsafe.Pointer]interface{}{}
	mutex sync.Mutex
)

func Save(v interface{}) unsafe.Pointer {
	if v == nil {
		return nil
	}
	var ptr unsafe.Pointer
	mutex.Lock()
	defer mutex.Unlock()
	// Generate real fake C pointer.
	// This pointer will not store any data, but will bi used for indexing purposes.
	// Since Go doest allow to cast dangling pointer to unsafe.Pointer, we do rally allocate one byte.
	// Why we need indexin, because Go doest allow C code to store pointers to Go data.
	ptr = C.malloc(C.size_t(1))
	if ptr == nil {
		mutex.Unlock()
		panic("can't allocate 'cgo-pointer hack index pointer': ptr == nil")
	}
	store[ptr] = v
	return ptr
}

func Restore(ptr unsafe.Pointer) interface{} {
	if ptr == nil {
		return nil
	}
	mutex.Lock()
	defer mutex.Unlock()
	if v, ok := store[ptr]; ok {
		return v
	}
	return nil
}

func Unref(ptr unsafe.Pointer) {
	if ptr == nil {
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	delete(store, ptr)
	C.free(ptr)
}
