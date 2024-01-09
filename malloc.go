package malloc

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

const Libc = "libc.so.6"

var malloc func(size uint) unsafe.Pointer
var free func(pointer unsafe.Pointer)

func init() {
	libc, err := purego.Dlopen(Libc, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	purego.RegisterLibFunc(&malloc, libc, "malloc")
	purego.RegisterLibFunc(&free, libc, "free")
}

func Malloc[T any](arena ...*Arena) *T {
	ptr := malloc(size[T]())
	if len(arena) > 0 {
		arena[0].register(ptr)
	}
	return (*T)(ptr)
}

func MallocSlice[T any](n uint, arena ...*Arena) []T {
	ptr := malloc(size[T]() * n)
	if len(arena) > 0 {
		arena[0].register(ptr)
	}
	return unsafe.Slice((*T)(ptr), n)
}

func Free[T any](t *T) {
	free(unsafe.Pointer(t))
}

func FreeSlice[T any](slice []T) {
	free(unsafe.Pointer(&slice[0]))
}

func size[T any]() uint {
	var t T
	return uint(unsafe.Sizeof(t))
}
