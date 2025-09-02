package cli

import (
	"fmt"

	"github.com/stevexciv/golang-todo-cli/tasks"
)

type DeleteCommand struct {
	Id int
}

func (d *DeleteCommand) Execute(m tasks.Manager) (string, error) {
	err := m.DeleteTask(d.Id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Deleted task (ID: %d)", d.Id), nil
}
