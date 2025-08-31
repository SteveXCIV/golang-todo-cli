package cli

import "github.com/stevexciv/golang-todo-cli/tasks"

type CompleteCommand struct{}

func (c *CompleteCommand) Execute(m *tasks.Manager) (string, error) {
	panic("not implemented!")
}
