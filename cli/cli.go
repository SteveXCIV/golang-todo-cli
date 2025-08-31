package cli

import (
	"github.com/stevexciv/golang-todo-cli/tasks"
)

type Command interface {
	Execute(m *tasks.Manager) (string, error)
}

func Parse(args *[]string) (Command, error) {
	panic("not implemented!")
}
