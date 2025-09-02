package cli

import (
	"fmt"
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

func (a *AddCommand) Execute(m tasks.Manager) (string, error) {
	added, err := m.AddTask(a.Title, a.Priority, func(t time.Time) tasks.DueDate {
		return a.Due.IntoDueDate(t)
	}, a.Category)

	if err != nil {
		return "", err
	}
	return "Added task '" + added.Title + "' successfully (ID: " + fmt.Sprint(added.Id) + ")", nil
}

func (d *DueToday) IntoDueDate(t time.Time) tasks.DueDate {
	return tasks.DueDate(t)
}

func (d *DueTomorrow) IntoDueDate(t time.Time) tasks.DueDate {
	return tasks.DueDate(t.Add(time.Hour * 24))
}

func (d *DueInDays) IntoDueDate(t time.Time) tasks.DueDate {
	return tasks.DueDate(t.Add(time.Hour * 24 * time.Duration(d.Days)))
}

func (d *DueOnDate) IntoDueDate(t time.Time) tasks.DueDate {
	return tasks.DueDate(d.At)
}
