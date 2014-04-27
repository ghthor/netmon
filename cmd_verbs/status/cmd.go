package status

import (
	"fmt"
)

type cmd struct {
}

func (c cmd) Exec([]string) error {
	fmt.Println("Executing the command bound to `status` verb")
	return nil
}

func (c cmd) Summary() string {
	return "how to use `status` verb"
}

var Cmd = &cmd{}
