package main

import(
	"fmt"
	"./decorator"
)

func main(){
	a := decorator.ConcreteComponent{"One Concr Comp"}
	b := decorator.BaseDecorator{a}
	c := decorator.BaseDecorator{b.Wrappee}
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
