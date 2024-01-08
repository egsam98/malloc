package malloc

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

const Libc = "libc.so.6"

var malloc func(size uint) unsafe.Pointer
var free func(pointer unsafe.Pointer)

// TODO calloc

func init() {
	libc, err := purego.Dlopen(Libc, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	purego.RegisterLibFunc(&malloc, libc, "malloc")
	purego.RegisterLibFunc(&free, libc, "free")
}

func Malloc[T any]() *T {
	return (*T)(malloc(size[T]()))
}

func Free[T any](t *T) {
	free(unsafe.Pointer(t))
}

func size[T any]() uint {
	var t T
	return uint(unsafe.Sizeof(t))
}
