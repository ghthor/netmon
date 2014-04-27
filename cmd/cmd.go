package cmd

import (
	"errors"
	"fmt"
)

// An interface to an executable command
type Cmd interface {
	Exec(args []string) error
}

// A map of verbs -> Command interfaces
type Catalog map[string]Cmd

func NewCatalog() Catalog {
	return make(map[string]Cmd)
}

// Register a verb to a Command
func (c Catalog) Register(verb string, cmd Cmd) error {
	if v, exists := c[verb]; exists {
		return errors.New(fmt.Sprintf("verb already registered: %s", v))
	}

	c[verb] = cmd
	return nil
}

// Perform a catalog[verb] key value lookup
func (c Catalog) MatchVerb(verb string) Cmd {
	return c[verb]
}

func (c Catalog) MatchExists(verb string) (Cmd, bool) {
	cmd, exists := c[verb]
	return cmd, exists
}

// A basic implementation of the Cmd interface
type C struct{}

func (c *C) Exec(args []string) error {
	return nil
}