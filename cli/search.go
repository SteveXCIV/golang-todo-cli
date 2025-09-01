package cli

import "github.com/stevexciv/golang-todo-cli/tasks"

type SearchCommand struct {
	Query string
}

func (s *SearchCommand) Execute(m *tasks.Manager) (string, error) {
	panic("not implemented")
}
