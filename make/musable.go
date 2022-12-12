//go:build ignore

package main

import (
	"reflect"

	"github.com/ymz-ncnk/musgo"
	mgi "github.com/ymz-ncnk/musgo_uint32"
)

func main() {
	musGo, err := musgo.New()
	if err != nil {
		panic(err)
	}
	unsafe := false
	err = musGo.Generate(reflect.TypeOf((*mgi.Uint32)(nil)).Elem(), unsafe)
	if err != nil {
		panic(err)
	}
}
