package decorator

type ConcreteDecorator struct {
	BaseDecorator
	link func() string
}

func (c *ConcreteDecorator) Extra() string{
	return "EXTRA"
}
