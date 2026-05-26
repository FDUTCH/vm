package vm

import "unsafe"

func cast[T any, V any](ptr *T) *V {
	return (*V)(unsafe.Pointer(ptr))
}
