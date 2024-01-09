package malloc

import (
	"testing"
)

func TestName(t *testing.T) {
	var arena Arena
	v1 := Malloc[byte](&arena)
	*v1 = 'd'
	v2 := Malloc[byte](&arena)
	*v2 = 'a'
	arena.Free()
	t.Log(string(*v1), string(*v2))
}
