package cli

import (
	"github.com/stevexciv/golang-todo-cli/tasks"
)

type SearchCommand struct {
	Query string
}

func (s *SearchCommand) Execute(m tasks.Manager) (string, error) {
	t, err := m.SearchTasks(s.Query)
	if err != nil {
		return "", err
	}
	return tasks.RenderTable(t), nil
}
