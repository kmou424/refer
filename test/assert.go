package test

import (
	"fmt"
	"testing"
)

func MustEqual[T comparable](t *testing.T, a, b T) {
	if a != b {
		err := fmt.Errorf("%v != %v", a, b)
		t.Error(err)
		panic(err)
	}
}
