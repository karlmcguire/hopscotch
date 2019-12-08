package hopscotch

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := New(16)

	m.Set(5, []byte("five"))
	fmt.Println(m)
	m.Set(21, []byte("twenty one"))
	fmt.Println(m)
	m.Set(37, []byte("thirty seven"))
	fmt.Println(m)
}
