package malloc

import "testing"

func TestName(t *testing.T) {
	v := Malloc[string]()
	*v = "test"
	t.Log(*v)
	Free(v)
	t.Log(*v) // panics
}
