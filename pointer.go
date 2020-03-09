package pointer

// #include <stdlib.h>
import "C"
import (
	"sync"
	"unsafe"
)

var (
	mutex sync.Mutex
	store = map[unsafe.Pointer]interface{}{}
)

func Save(v interface{}) unsafe.Pointer {
	if v == nil {
		return nil
	}

	// Generates a real fake C pointer.
	// The pointer won't store any data but be used for indexing purposes.
	// As Go doesn't allow to cast a dangling pointer to "unsafe.Pointer", we do really allocate one byte.
	// Indexing is needed because Go doesn't allow C code to store pointers to Go data.
	var ptr unsafe.Pointer = C.malloc(C.size_t(1))
	if ptr == nil {
		panic("Can't allocate 'cgo-pointer hack index pointer': ptr == nil")
	}

	mutex.Lock()
	store[ptr] = v
	mutex.Unlock()

	return ptr
}

func Restore(ptr unsafe.Pointer) (v interface{}) {
	if ptr == nil {
		return nil
	}

	mutex.Lock()
	v = store[ptr]
	mutex.Unlock()
	return
}

func Unref(ptr unsafe.Pointer) {
	if ptr == nil {
		return
	}

	mutex.Lock()
	delete(store, ptr)
	mutex.Unlock()

	C.free(ptr)
}
