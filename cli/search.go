package cli

import (
	"strings"

	"github.com/stevexciv/golang-todo-cli/tasks"
)

type SearchCommand struct {
	Query string
}

func (s *SearchCommand) Execute(m tasks.Manager) (string, error) {
	tasks, err := m.SearchTasks(s.Query)
	if err != nil {
		return "", err
	}
	builder := strings.Builder{}
	for i, task := range tasks {
		builder.WriteString(task.Title)
		if i < len(tasks)-1 {
			builder.WriteString("\n")
		}
	}
	return builder.String(), nil
}
