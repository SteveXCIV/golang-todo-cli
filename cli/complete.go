package cli

import (
	"fmt"

	"github.com/stevexciv/golang-todo-cli/tasks"
)

type CompleteCommand struct {
	Id int
}

func (c *CompleteCommand) Execute(m tasks.Manager) (string, error) {
	err := m.CompleteTask(c.Id)
	if err != nil {
		return "", err
	}
	return "Task completed successfully (ID: " + fmt.Sprint(c.Id) + ")", nil
}
