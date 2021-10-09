package constrant_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Readable, Writable, Executable)
}
