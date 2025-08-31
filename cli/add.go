package cli

import "github.com/stevexciv/golang-todo-cli/tasks"

type AddCommand struct{}

func (a *AddCommand) Execute(m *tasks.Manager) (string, error) {
	panic("not implemented")
}
