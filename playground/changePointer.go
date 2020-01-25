package playground

import "fmt"

type Foo struct {
	name string
}

func (foo *Foo) changeAddr() {
	fmt.Printf("struct passed %p, %p\n", foo, &foo)
	foo = &Foo{}
	fmt.Printf("struct changed %p, %p\n", foo, &foo)
}

func changeAddr(foo *Foo) {
	fmt.Printf("param passed %p, %p\n", foo, &foo)
	foo = &Foo{}
	fmt.Printf("param changed %p, %p\n", foo, &foo)
}

// func main() {
// 	a := &Foo{name: "a"}
// 	fmt.Printf("a1 %p, %p\n", a, &a)
// 	a.changeAddr()
// 	fmt.Printf("a2 %p, %p\n", a, &a)
// 	changeAddr(a)
// 	fmt.Printf("a3 %p, %p\n", a, &a)
// }
