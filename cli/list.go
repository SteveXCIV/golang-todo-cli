package cli

import "github.com/stevexciv/golang-todo-cli/tasks"

type ListCommand struct{}

func (l *ListCommand) Execute(m *tasks.Manager) (string, error) {
	panic("not implemented!")
}
