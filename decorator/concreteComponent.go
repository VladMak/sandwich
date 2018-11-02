package decorator

type ConcreteComponent struct {
	Name string
}

func (c ConcreteComponent) Execute() string {
	return "From Concret Component" + c.Name
}
