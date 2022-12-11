package musgo_uint32

import "testing"

func TestMusgoUint32(t *testing.T) {
	var n uint32 = 5
	buf := make([]byte, Uint32(n).SizeMUS())
	Uint32(n).MarshalMUS(buf)

	var an uint32
	(*Uint32)(&an).UnmarshalMUS(buf)

	if n != an {
		t.Fatal("something went wrong")
	}
}
