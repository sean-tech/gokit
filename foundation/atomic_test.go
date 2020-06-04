package foundation

import (
	"testing"
	"time"
)

func TestInt32(t *testing.T) {
	atom := NewInt32(42)
	atom.Store(42)

}

func TestInt64(t *testing.T) {
	atom := NewInt64(42)
	atom.Store(42)

}

func TestUint32(t *testing.T) {
	atom := NewUint32(42)

	atom.Store(42)
}

func TestUint64(t *testing.T) {
	atom := NewUint64(42)

	atom.Store(42)
}

func TestBool(t *testing.T) {
	atom := NewBool(false)
	atom.Store(false)
	//prev := atom.Swap(false)
	//prev = atom.Swap(true)
}

func TestFloat64(t *testing.T) {
	atom := NewFloat64(4.2)
	atom.Store(42.0)
}

func TestDuration(t *testing.T) {
	atom := NewDuration(5 * time.Minute)
	atom.Store(10 * time.Minute)
}

func TestValue(t *testing.T) {
	var v Value

	v.Store(42)
	v.Store(84)
}
