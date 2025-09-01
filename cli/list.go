package cli

import "github.com/stevexciv/golang-todo-cli/tasks"

type ListCommand struct {
	StatusFilter   *tasks.Status
	PriorityFilter *tasks.Priority
	CategoryFilter string
	OverdueFilter  bool
}

func (l *ListCommand) Execute(m *tasks.Manager) (string, error) {
	panic("not implemented!")
}
