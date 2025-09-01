package cli

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stevexciv/golang-todo-cli/tasks"
)

func TestParseInvalidTableDriven(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "empty",
			args: []string{},
		},
		{
			name: "invalid command single",
			args: []string{"foobar"},
		},
		{
			name: "invalid command multi",
			args: []string{"foo", "bar", "fizz", "buzz"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd, err := Parse(&tt.args)

			if err == nil {
				t.Fatalf("expected err, got: %v", cmd)
			}
			if !strings.Contains(strings.ToLower(err.Error()), "invalid") {
				t.Fatalf("expected 'invlaid' in err, got: %v", err)
			}
		})
	}
}

func TestParseAddTableDriven(t *testing.T) {
	invalid := []struct {
		name   string
		args   []string
		errMsg string
	}{
		{
			name:   "add empty title",
			args:   []string{"add", ""},
			errMsg: "title cannot be empty",
		},
		{
			name:   "add invalid priority",
			args:   []string{"add", "Call dentist", "--priority", "mega-high"},
			errMsg: "invalid priority format",
		},
		{
			name:   "add invalid date",
			args:   []string{"add", "Call dentist", "--due", "foobar"},
			errMsg: "invalid date format",
		},
	}
	tests := []struct {
		name   string
		args   []string
		addCmd AddCommand
	}{
		{
			name: "add minimal",
			args: []string{"add", "Call dentist"},
			addCmd: AddCommand{
				Title:    "Call dentist",
				Priority: tasks.Medium,
				Due:      &DueToday{},
			},
		},
		{
			name: "add with priority",
			args: []string{"add", "Call dentist", "--priority", "high"},
			addCmd: AddCommand{
				Title:    "Call dentist",
				Priority: tasks.High,
				Due:      &DueToday{},
			},
		},
		{
			name: "add with explicit today",
			args: []string{"add", "Call dentist", "--due", "today"},
			addCmd: AddCommand{
				Title:    "Call dentist",
				Priority: tasks.Medium,
				Due:      &DueToday{},
			},
		},
		{
			name: "add with tomorrow due date",
			args: []string{"add", "Call dentist", "--due", "tomorrow"},
			addCmd: AddCommand{
				Title:    "Call dentist",
				Priority: tasks.Medium,
				Due:      &DueTomorrow{},
			},
		},
		{
			name: "add with relative due date",
			args: []string{"add", "Call dentist", "--due", "+7d"},
			addCmd: AddCommand{
				Title:    "Call dentist",
				Priority: tasks.Medium,
				Due:      &DueInDays{Days: 7},
			},
		},
		{
			name: "add with absolute due date",
			args: []string{"add", "Call dentist", "--due", "2020-04-10"},
			addCmd: AddCommand{
				Title:    "Call dentist",
				Priority: tasks.Medium,
				Due:      &DueOnDate{At: time.Date(2020, 4, 10, 0, 0, 0, 0, time.UTC)},
			},
		},
	}

	for _, it := range invalid {
		t.Run(it.name, func(t *testing.T) {
			cmd, err := Parse(&it.args)

			if err == nil {
				t.Fatalf("expected err, got: %v", cmd)
			}
			if !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(it.errMsg)) {
				t.Fatalf("unexpected err text: wanted='%v', got=%v", it.errMsg, err)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd, err := Parse(&tt.args)

			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			addCmd, ok := cmd.(*AddCommand)
			if !ok {
				t.Fatalf("expected add command, got: %v", cmd)
			}
			if !reflect.DeepEqual(&tt.addCmd, addCmd) {
				t.Fatalf("unexpected command: wanted=%v, got=%v", tt.addCmd, addCmd)
			}
		})
	}
}

func TestParseListTableDriven(t *testing.T) {
	invalid := []struct {
		name   string
		args   []string
		errMsg string
	}{
		{
			name:   "list invalid priority",
			args:   []string{"list", "--priority", "mega-high"},
			errMsg: "invalid priority format",
		},
		{
			name:   "list invalid status",
			args:   []string{"list", "--status", "maybe"},
			errMsg: "invalid status format",
		},
		{
			name:   "list missing priority value",
			args:   []string{"list", "--priority"},
			errMsg: "priority flag requires a value",
		},
		{
			name:   "list missing status value",
			args:   []string{"list", "--status"},
			errMsg: "status flag requires a value",
		},
		{
			name:   "list missing category value",
			args:   []string{"list", "--category"},
			errMsg: "category flag requires a value",
		},
	}
	tests := []struct {
		name    string
		args    []string
		listCmd ListCommand
	}{
		{
			name: "list basic",
			args: []string{"list"},
			listCmd: ListCommand{
				StatusFilter:   nil,
				PriorityFilter: nil,
				CategoryFilter: "",
				OverdueFilter:  false,
			},
		},
		{
			name: "list with status filter",
			args: []string{"list", "--status", "pending"},
			listCmd: ListCommand{
				StatusFilter:   &[]tasks.Status{tasks.Pending}[0],
				PriorityFilter: nil,
				CategoryFilter: "",
				OverdueFilter:  false,
			},
		},
		{
			name: "list with priority filter",
			args: []string{"list", "--priority", "high"},
			listCmd: ListCommand{
				StatusFilter:   nil,
				PriorityFilter: &[]tasks.Priority{tasks.High}[0],
				CategoryFilter: "",
				OverdueFilter:  false,
			},
		},
		{
			name: "list with category filter",
			args: []string{"list", "--category", "personal"},
			listCmd: ListCommand{
				StatusFilter:   nil,
				PriorityFilter: nil,
				CategoryFilter: "personal",
				OverdueFilter:  false,
			},
		},
		{
			name: "list with overdue filter",
			args: []string{"list", "--overdue"},
			listCmd: ListCommand{
				StatusFilter:   nil,
				PriorityFilter: nil,
				CategoryFilter: "",
				OverdueFilter:  true,
			},
		},
		{
			name: "list with multiple filters",
			args: []string{"list", "--status", "pending", "--priority", "high"},
			listCmd: ListCommand{
				StatusFilter:   &[]tasks.Status{tasks.Pending}[0],
				PriorityFilter: &[]tasks.Priority{tasks.High}[0],
				CategoryFilter: "",
				OverdueFilter:  false,
			},
		},
		{
			name: "list with all filters",
			args: []string{"list", "--status", "completed", "--priority", "low", "--category", "work", "--overdue"},
			listCmd: ListCommand{
				StatusFilter:   &[]tasks.Status{tasks.Completed}[0],
				PriorityFilter: &[]tasks.Priority{tasks.Low}[0],
				CategoryFilter: "work",
				OverdueFilter:  true,
			},
		},
	}

	for _, it := range invalid {
		t.Run(it.name, func(t *testing.T) {
			cmd, err := Parse(&it.args)

			if err == nil {
				t.Fatalf("expected err, got: %v", cmd)
			}
			if !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(it.errMsg)) {
				t.Fatalf("unexpected err text: wanted='%v', got=%v", it.errMsg, err)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd, err := Parse(&tt.args)

			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			listCmd, ok := cmd.(*ListCommand)
			if !ok {
				t.Fatalf("expected list command, got: %v", cmd)
			}
			if !reflect.DeepEqual(&tt.listCmd, listCmd) {
				t.Fatalf("unexpected command: wanted=%v, got=%v", tt.listCmd, listCmd)
			}
		})
	}
}

func TestParseSearchListTableDriven(t *testing.T) {}

func TestParseCompleteTableDriven(t *testing.T) {}

func TestParseDeleteTableDriven(t *testing.T) {}
