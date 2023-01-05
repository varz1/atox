package atox

import (
	"testing"
)

func TestCommon(t *testing.T) {
	MustN[int64]("1")

	F(1)
}

type Uint8 uint8

var eq = func(x, y any, t *testing.T) {
	if x != y {
		t.Errorf("%v (%T) != %v (%T)", x, x, y, y)
	}
}

func TestN(t *testing.T) {
	eq(MustN[float32]("1.0"), float32(1.0), t)
	eq(MustN[float64]("1.0"), float64(1.0), t)
	eq(MustN[uint8]("1"), uint8(1), t)
	eq(MustN[Uint8]("1"), Uint8(1), t)
}

func TestF(t *testing.T) {
	eq(F(1), "1", t)
	eq(F(1.0), "1", t)
	eq(F(111), "111", t)
	eq(F("xx"), "xx", t)
}
