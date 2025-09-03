package cli

import (
	"github.com/stevexciv/golang-todo-cli/tasks"
)

type ListCommand struct {
	StatusFilter   *tasks.Status
	PriorityFilter *tasks.Priority
	CategoryFilter string
	OverdueFilter  bool
}

func (l *ListCommand) Execute(m tasks.Manager) (string, error) {
	t, err := m.ListTasks(l.StatusFilter, l.PriorityFilter, l.CategoryFilter, l.OverdueFilter)
	if err != nil {
		return "", err
	}
	return tasks.RenderTable(t), nil
}
