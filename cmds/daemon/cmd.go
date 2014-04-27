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

var Cmd = &cmd{}
