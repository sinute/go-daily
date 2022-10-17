package main

import (
	"fmt"
)

func main() {
	f := Foo{p: new(nested), v: nested{}}
	fmt.Printf("foo.p: %p\n", &f.p.value)
	fmt.Printf("foo.v: %p\n", &f.v.value)
	f.PointerReceiver()
	f.ValueReceiver()
}

type Foo struct {
	p *nested
	v nested
}

type nested struct {
	value int
}

func (f *Foo) PointerReceiver() {
	fmt.Printf("pointerReceiver.p: %p\n", &f.p.value)
	fmt.Printf("pointerReceiver.v: %p\n", &f.v.value)
}

func (f Foo) ValueReceiver() {
	fmt.Printf("valueReceiver.p: %p\n", &f.p.value)
	fmt.Printf("valueReceiver.v: %p\n", &f.v.value)
}
