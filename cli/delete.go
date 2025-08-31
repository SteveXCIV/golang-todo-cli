package cli

import (
	"github.com/stevexciv/golang-todo-cli/tasks"
)

type DeleteCommand struct{}

func (d *DeleteCommand) Execute(m *tasks.Manager) (string, error) {
	panic("not implemented!")
}
