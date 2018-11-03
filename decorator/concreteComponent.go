package decorator

import(
	"reflect"
)

//In the class adding new function from Decorator
type ConcreteComponent struct {
	Name string
}

func (c ConcreteComponent) Execute() reflect.Type {
	return reflect.TypeOf(c)
}
