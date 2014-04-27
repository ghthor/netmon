package daemon

import (
	"fmt"
)

type cmd struct {
}

func (c cmd) Exec([]string) error {
	fmt.Println("Executing the command bound to `daemon` verb")
	return nil
}

func (c cmd) Summary() string {
	return "how to use `daemon` verb"
}

var Cmd = &cmd{}
