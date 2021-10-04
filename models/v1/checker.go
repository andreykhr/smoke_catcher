package v1

type Checker struct {
	name string
}

func (c *Checker) Init (name string) {
	c.name = name
}

func (c *Checker) Check() (string, error) {
	return c.name, nil
	//return c.name, errors.New("some error")
}