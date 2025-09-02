package cli

import (
	"fmt"
	"strings"

	"github.com/stevexciv/golang-todo-cli/tasks"
)

type ListCommand struct {
	StatusFilter   *tasks.Status
	PriorityFilter *tasks.Priority
	CategoryFilter string
	OverdueFilter  bool
}

func (l *ListCommand) Execute(m tasks.Manager) (string, error) {
	tasks, err := m.ListTasks(l.StatusFilter, l.PriorityFilter, l.CategoryFilter, l.OverdueFilter)
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	for _, t := range tasks {
		sb.WriteString(fmt.Sprint(t) + "\n")
	}
	return sb.String(), nil
}
