package cli

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestCompleteExecuteTableDriven(t *testing.T) {
	m := newMockManager(testTime)
	tests := []struct {
		name        string
		completeCmd CompleteCommand
		wantCall    completeCall
		want        string
	}{
		{
			name: "complete task",
			completeCmd: CompleteCommand{
				Id: 1,
			},
			wantCall: completeCall{
				id: 1,
			},
			want: "Task 1 completed successfully",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.reset()

			got, err := tt.completeCmd.Execute(&m)

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if !slices.Contains(m.completeCalls, tt.wantCall) {
				t.Fatalf("missing expected call: wanted=%v, got=%v", tt.wantCall, m.completeCalls)
			}
			got = strings.ToLower(got)
			want := strings.ToLower(tt.want)
			if !strings.Contains(got, want) {
				t.Fatalf("expected output to contain '%v', got: %v", want, got)
			}
		})
	}
}

func TestCompleteExecuteErrorsTableDriven(t *testing.T) {
	m := newMockManager(testTime)
	tests := []struct {
		name        string
		completeCmd CompleteCommand
		mockErr     error
		wantErr     string
	}{
		{
			name:        "task not found",
			completeCmd: CompleteCommand{Id: 999},
			mockErr:     errors.New("task not found"),
			wantErr:     "task not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.reset()

			_, err := tt.completeCmd.Execute(&m)

			if err == nil {
				t.Fatalf("expected error, got none")
			}
			if !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(tt.wantErr)) {
				t.Fatalf("expected error to contain '%v', got: %v", tt.wantErr, err)
			}
		})
	}
}
