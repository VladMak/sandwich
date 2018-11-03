package decorator

type ConcreteDecorator2 struct {
	BaseDecorator
}

func (c *ConcreteDecorator2) TestFunc() string{
	return "test"
}
