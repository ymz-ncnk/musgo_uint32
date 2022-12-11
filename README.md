# MusGo_uint32
Provides serialization of the Golang's `uint32` type.

# How to use
Simply cast `uint32` to `musgo_uint32.Uint32`:
```go
	var n uint32 = 5
	// Marshal
	buf := make([]byte, musgo_uint32.Uint32(n).SizeMUS())
	musgo_uint32.Uint32(n).MarshalMUS(buf)
	// Unmarshal
	_, err := (*musgo_uint32.Uint32)(&n).UnmarshalMUS(buf)
	if err != nil {
		panic(err)
	}
```

# More info
You can find at [github.com/ymz-ncnk/musgo](https://github.com/ymz-ncnk/musgo).

