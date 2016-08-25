package example

/*
void *pass_pointer(void* p) {
  return p;
}
*/
import "C"
import "unsafe"

func PassPointer(ptr unsafe.Pointer) unsafe.Pointer {
	return C.pass_pointer(ptr)
}
