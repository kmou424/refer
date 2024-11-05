package test

import (
	"github.com/kmou424/refer"
	"testing"
)

func TestNamespace(t *testing.T) {
	type st struct {
		A int
		B string
	}

	var s = &st{
		A: 1,
		B: "hello",
	}

	refer.Namespace(refer.NSPkg)
	refer.Bind(s)

	refer.Namespace(refer.NSGlobal)
	MustEqual(t, refer.Ref[st](), nil)

	refer.Namespace(refer.NSPkg)
	MustEqual(t, refer.Ref[st](), s)

	MustEqual(t, refer.Ref[st]().A, 1)
	MustEqual(t, refer.Ref[st]().B, "hello")
}
