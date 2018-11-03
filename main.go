package main

import(
	"fmt"
	"./decorator"
)

func main(){
	a := decorator.ConcreteComponent{"One Concr Comp"}
	b := new(decorator.ConcreteDecorator)
	b.Wrappee = a
	c := new(decorator.ConcreteDecorator2)
	c.Wrappee = b.Wrappee
	fmt.Println(a.Execute())
	fmt.Println(b.Wrappee.Execute())
	fmt.Println(c.Wrappee.Execute())
	fmt.Println(b.Extra())
	fmt.Println(c.TestFunc())
}
