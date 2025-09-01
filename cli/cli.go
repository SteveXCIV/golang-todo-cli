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
	// handle simple args case
	args := *a
	if len(args) == 0 {
		return nil, fmt.Errorf("invalid command: must specify one of 'add', 'list', 'search', 'complete', or 'delete'")
	}

	switch args[0] {
	case "add":
		return parseAddCmd(args[1:])
	case "list":
		return parseListCmd(args[1:])
	case "search":
		return parseSearchCmd(args[1:])
	case "complete":
		return parseCompleteCmd(args[1:])
	case "delete":
		return parseDeleteCmd(args[1:])
	}
	return nil, fmt.Errorf("invalid command: must specify one of 'add', 'list', 'search', 'complete', or 'delete'")
}

func parseAddCmd(a []string) (*AddCommand, error) {
	if len(a) == 0 {
		return nil, fmt.Errorf("add error: title cannot be empty")
	}
	title := a[0]
	if strings.TrimSpace(title) == "" {
		return nil, fmt.Errorf("add error: title cannot be empty")
	}

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
		match := plusDaysRegex.FindStringSubmatch(*addDue)
		days, err := strconv.Atoi(match[1])
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

func parseListCmd(a []string) (*ListCommand, error) {
	listFlagSet := flag.NewFlagSet("list", flag.ExitOnError)
	listPriority := listFlagSet.String("priority", "", "Priority filter: low, medium, high")
	listStatus := listFlagSet.String("status", "", "Status filter: pending, completed")
	listCategory := listFlagSet.String("category", "", "Category filter")
	listOverdue := listFlagSet.Bool("overdue", false, "Show only overdue tasks")

	if err := listFlagSet.Parse(a); err != nil {
		return nil, err
	}

	var priorityFilter *tasks.Priority
	if *listPriority != "" {
		switch strings.ToLower(*listPriority) {
		case "low":
			p := tasks.Low
			priorityFilter = &p
		case "medium":
			p := tasks.Medium
			priorityFilter = &p
		case "high":
			p := tasks.High
			priorityFilter = &p
		default:
			return nil, fmt.Errorf("invalid priority format: must be 'low', 'medium', or 'high'")
		}
	}

	var statusFilter *tasks.Status
	if *listStatus != "" {
		switch strings.ToLower(*listStatus) {
		case "pending":
			s := tasks.Pending
			statusFilter = &s
		case "completed":
			s := tasks.Completed
			statusFilter = &s
		default:
			return nil, fmt.Errorf("invalid status format: must be 'pending' or 'completed'")
		}
	}

	listCmd := &ListCommand{
		StatusFilter:   statusFilter,
		PriorityFilter: priorityFilter,
		CategoryFilter: *listCategory,
		OverdueFilter:  *listOverdue,
	}

	return listCmd, nil
}

func parseSearchCmd(a []string) (*SearchCommand, error) {
	panic("not implemented")
}

func parseCompleteCmd(a []string) (*CompleteCommand, error) {
	panic("not implemented")
}

func parseDeleteCmd(a []string) (*DeleteCommand, error) {
	panic("not implemented")
}
