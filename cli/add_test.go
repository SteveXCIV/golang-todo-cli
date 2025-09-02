package cli

import (
	"errors"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/stevexciv/golang-todo-cli/tasks"
)

var testTime = time.Date(2024, 4, 10, 0, 0, 0, 0, time.UTC)

func TestAddExecuteTableDriven(t *testing.T) {
	m := newMockManager(testTime)
	tests := []struct {
		name       string
		addCmd     AddCommand
		mockReturn *tasks.Task
		wantCall   addCall
		want       string
	}{
		{
			name: "add due today",
			addCmd: AddCommand{
				Title:    "Call dentist",
				Priority: tasks.High,
				Due:      &DueToday{},
			},
			mockReturn: &tasks.Task{
				Id:       1,
				Title:    "Call dentist",
				Priority: tasks.High,
				DueDate:  tasks.DueDate(testTime),
				Status:   tasks.Pending,
			},
			wantCall: addCall{
				title:    "Call dentist",
				priority: tasks.High,
				dueDate:  tasks.DueDate(testTime),
			},
			want: "Added task 'Call dentist' successfully",
		},
		{
			name: "add due tomorrow",
			addCmd: AddCommand{
				Title:    "Call dentist",
				Priority: tasks.Medium,
				Due:      &DueTomorrow{},
			},
			mockReturn: &tasks.Task{
				Id:       1,
				Title:    "Call dentist",
				Priority: tasks.Medium,
				DueDate:  tasks.DueDate(testTime.Add(time.Hour * 24)),
				Status:   tasks.Pending,
			},
			wantCall: addCall{
				title:    "Call dentist",
				priority: tasks.Medium,
				dueDate:  tasks.DueDate(testTime.Add(time.Hour * 24)),
			},
			want: "Added task 'Call dentist' successfully",
		},
		{
			name: "add due relative",
			addCmd: AddCommand{
				Title:    "Call dentist",
				Priority: tasks.Medium,
				Due:      &DueInDays{Days: 7},
			},
			mockReturn: &tasks.Task{
				Id:       1,
				Title:    "Call dentist",
				Priority: tasks.Medium,
				DueDate:  tasks.DueDate(testTime.Add(time.Hour * 24 * 7)),
				Status:   tasks.Pending,
			},
			wantCall: addCall{
				title:    "Call dentist",
				priority: tasks.Medium,
				dueDate:  tasks.DueDate(testTime.Add(time.Hour * 24 * 7)),
			},
			want: "Added task 'Call dentist' successfully",
		},
		{
			name: "add due absolute",
			addCmd: AddCommand{
				Title:    "Call dentist",
				Priority: tasks.Medium,
				Due:      &DueOnDate{At: time.Date(2024, 4, 31, 0, 0, 0, 0, time.UTC)},
			},
			mockReturn: &tasks.Task{
				Id:       1,
				Title:    "Call dentist",
				Priority: tasks.Medium,
				DueDate:  tasks.DueDate(time.Date(2024, 4, 31, 0, 0, 0, 0, time.UTC)),
				Status:   tasks.Pending,
			},
			wantCall: addCall{
				title:    "Call dentist",
				priority: tasks.Medium,
				dueDate:  tasks.DueDate(time.Date(2024, 4, 31, 0, 0, 0, 0, time.UTC)),
			},
			want: "Added task 'Call dentist' successfully",
		},
		{
			name: "add with category",
			addCmd: AddCommand{
				Title:    "Call dentist",
				Priority: tasks.Medium,
				Due:      &DueOnDate{At: time.Date(2024, 4, 31, 0, 0, 0, 0, time.UTC)},
				Category: "Health",
			},
			mockReturn: &tasks.Task{
				Id:       1,
				Title:    "Call dentist",
				Priority: tasks.Medium,
				DueDate:  tasks.DueDate(time.Date(2024, 4, 31, 0, 0, 0, 0, time.UTC)),
				Status:   tasks.Pending,
				Category: "Health",
			},
			wantCall: addCall{
				title:    "Call dentist",
				priority: tasks.Medium,
				dueDate:  tasks.DueDate(time.Date(2024, 4, 31, 0, 0, 0, 0, time.UTC)),
				category: "Health",
			},
			want: "Added task 'Call dentist' successfully",
		},
	}
	var manager tasks.Manager = &m

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.reset()
			m.addNextOk = tt.mockReturn

			// TODO: revisit this API to see if we can avoid double-pointer
			got, err := tt.addCmd.Execute(manager)

			if err != nil {
				t.Fatalf("exepcted no error, got %v", err)
			}
			if !slices.Contains(m.addCalls, tt.wantCall) {
				t.Fatalf("missing expected call: wanted=%v, got=%v", tt.wantCall, m.addCalls)
			}
			got = strings.ToLower(got)
			want := strings.ToLower(tt.want)
			if !strings.Contains(got, want) {
				t.Fatalf("expected output to contain '%v', got: %v", want, got)
			}
		})
	}
}

func TestAddExecuteErrorsTableDriven(t *testing.T) {
	m := newMockManager(testTime)
	tests := []struct {
		name    string
		addCmd  AddCommand
		mockErr error
		wantErr string
	}{
		{
			name: "add task error",
			addCmd: AddCommand{
				Title:    "Call dentist",
				Priority: tasks.Medium,
				Due:      &DueOnDate{At: time.Date(2024, 4, 31, 0, 0, 0, 0, time.UTC)},
			},
			mockErr: errors.New("failed to add task"),
			wantErr: "failed to add task",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.reset()
			m.addNextErr = tt.mockErr

			_, err := tt.addCmd.Execute(&m)

			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(tt.wantErr)) {
				t.Fatalf("expected error message to contain '%v', got: %v", tt.wantErr, err)
			}
		})
	}
}
