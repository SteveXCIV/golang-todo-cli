package cli

import (
	"time"

	"github.com/stevexciv/golang-todo-cli/tasks"
)

type IntoDueDate interface {
	IntoDueDate(time.Time) tasks.DueDate
}

type DueToday struct{}

type DueTomorrow struct{}

type DueInDays struct {
	Days int
}

type DueOnDate struct {
	At time.Time
}

type AddCommand struct {
	Title    string
	Priority tasks.Priority
	Due      IntoDueDate
	Category string
}

func (a *AddCommand) Execute(m *tasks.Manager) (string, error) {
	panic("not implemented")
}

func (d *DueToday) IntoDueDate(t time.Time) tasks.DueDate {
	panic("not implemented")
}

func (d *DueTomorrow) IntoDueDate(t time.Time) tasks.DueDate {
	panic("not implemented")
}

func (d *DueInDays) IntoDueDate(t time.Time) tasks.DueDate {
	panic("not implemented")
}

func (d *DueOnDate) IntoDueDate(t time.Time) tasks.DueDate {
	panic("not implemented")
}
