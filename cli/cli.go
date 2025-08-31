package cli

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/stevexciv/golang-todo-cli/tasks"
)

type Command interface {
	Execute(m *tasks.Manager) (string, error)
}

func Parse(a *[]string) (Command, error) {
	// common error messages
	const invalidCmdMsg = "invalid command: must specify one of 'add', 'list', 'search', 'complete', or 'delete'"

	// handle simple args case
	args := *a
	if len(args) == 0 {
		return nil, fmt.Errorf(invalidCmdMsg)
	}

	switch args[0] {
	case "add":
		return parseAddCmd(args[1:])
	}
	return nil, fmt.Errorf(invalidCmdMsg)
}

func parseAddCmd(a []string) (*AddCommand, error) {
	if len(a) == 0 {
		return nil, fmt.Errorf("add error: title cannot be empty")
	}
	title := a[0]

	addFlagSet := flag.NewFlagSet("add", flag.ExitOnError)
	addPriority := addFlagSet.String("priority", "medium", "Priority: low, medium (default), high")
	addDue := addFlagSet.String("due", "today", "Due date: today (default), tomorrow, +Xd (days), or yyyy-MM-dd")
	addCategory := addFlagSet.String("category", "", "Category: optional descriptive category")
	if err := addFlagSet.Parse(a[1:]); err != nil {
		return nil, err
	}

	var priority tasks.Priority
	switch strings.ToLower(*addPriority) {
	case "low":
		priority = tasks.Low
	case "medium":
		priority = tasks.Medium
	case "high":
		priority = tasks.High
	default:
		return nil, fmt.Errorf("invalid priority format: must be 'low', 'medium', or 'high'")
	}

	plusDaysRegex := regexp.MustCompile(`^\+(\d+)d$`)
	var due IntoDueDate
	switch {
	case *addDue == "today":
		due = &DueToday{}
	case *addDue == "tomorrow":
		due = &DueTomorrow{}
	case plusDaysRegex.MatchString(*addDue):
		match := plusDaysRegex.FindString(*addDue)
		days, err := strconv.Atoi(match)
		if err != nil {
			return nil, fmt.Errorf("invalid date format +Xd: %v", err)
		}
		due = &DueInDays{Days: days}
	default:
		at, err := time.Parse("2006-01-02", *addDue)
		if err != nil {
			return nil, fmt.Errorf("invalid date format yyyy-MM-dd: %v", err)
		}
		due = &DueOnDate{At: at}
	}

	category := strings.ToLower(*addCategory)
	addCmd := &AddCommand{
		Title:    title,
		Priority: priority,
		Due:      due,
		Category: category,
	}

	return addCmd, nil
}
