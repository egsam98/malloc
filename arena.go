package malloc

import "unsafe"

type Arena struct {
	ptrs []unsafe.Pointer // TODO check escape analysis
}

func (a *Arena) Free() {
	for _, ptr := range a.ptrs {
		free(ptr)
	}
}

func (a *Arena) register(ptr unsafe.Pointer) {
	a.ptrs = append(a.ptrs, ptr)
}
