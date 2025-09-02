package cli

import (
	"errors"
	"strings"
	"testing"

	"github.com/stevexciv/golang-todo-cli/tasks"
)

func TestListExecuteTableDriven(t *testing.T) {
	m := newMockManager(testTime)
	pending := tasks.Pending
	priority := tasks.High
	tests := []struct {
		name       string
		listCmd    ListCommand
		mockReturn []tasks.Task
		wantCall   listCall
		want       string
	}{
		{
			name: "list tasks",
			listCmd: ListCommand{
				StatusFilter:   &pending,
				PriorityFilter: &priority,
				CategoryFilter: "",
				OverdueFilter:  false,
			},
			mockReturn: []tasks.Task{{Id: 1, Title: "Task 1"}},
			wantCall: listCall{
				status:      &pending,
				priority:    &priority,
				category:    "",
				overdueOnly: false,
			},
			want: "Task 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.reset()
			m.listNextOk = tt.mockReturn

			got, err := tt.listCmd.Execute(&m)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if !strings.Contains(strings.ToLower(got), strings.ToLower(tt.want)) {
				t.Fatalf("expected response to contain '%v', got '%v'", tt.want, got)
			}
		})
	}
}

func TestListExecuteErrorsTableDriven(t *testing.T) {
	m := newMockManager(testTime)
	tests := []struct {
		name    string
		listCmd ListCommand
		mockErr error
		wantErr string
	}{
		{
			name: "list error",
			listCmd: ListCommand{
				StatusFilter:   nil,
				PriorityFilter: nil,
				CategoryFilter: "",
				OverdueFilter:  false,
			},
			mockErr: errors.New("list error"),
			wantErr: "list error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.reset()
			m.listNextErr = tt.mockErr

			_, err := tt.listCmd.Execute(&m)
			if err == nil {
				t.Fatalf("expected error, got none")
			}
			if !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(tt.wantErr)) {
				t.Fatalf("expected error to contain '%v', got '%v'", tt.wantErr, err)
			}
		})
	}
}
