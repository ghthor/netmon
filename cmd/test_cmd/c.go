package test_cmd

type Cmd struct {
	wasExecuted bool
}

func (c *Cmd) Exec([]string) error {
	c.wasExecuted = true
	return nil
}

var C *Cmd = &Cmd{}
