package cli

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestDeleteExecuteTableDriven(t *testing.T) {
	m := newMockManager(testTime)
	tests := []struct {
		name      string
		deleteCmd DeleteCommand
		wantCall  deleteCall
		want      string
	}{
		{
			name:      "delete existing item",
			deleteCmd: DeleteCommand{Id: 1},
			wantCall:  deleteCall{id: 1},
			want:      "deleted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.reset()

			ok, err := tt.deleteCmd.Execute(&m)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !slices.Contains(m.deleteCalls, tt.wantCall) {
				t.Errorf("wanted delete call: %v", tt.wantCall)
			}
			if !strings.Contains(strings.ToLower(ok), strings.ToLower(tt.want)) {
				t.Errorf("unexpected result: got '%v', want '%v'", ok, tt.want)
			}
		})
	}
}

func TestDeleteExecuteErrorsTableDriven(t *testing.T) {
	m := newMockManager(testTime)
	tests := []struct {
		name      string
		deleteCmd DeleteCommand
		deleteErr error
		wantErr   string
	}{
		{
			name:      "delete error",
			deleteCmd: DeleteCommand{Id: 1},
			deleteErr: errors.New("delete failed"),
			wantErr:   "delete failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.reset()

			_, err := tt.deleteCmd.Execute(&m)

			if err == nil {
				t.Fatalf("expected error, got none")
			}
			if !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(tt.wantErr)) {
				t.Errorf("unexpected error message: got %q, want %q", err.Error(), tt.wantErr)
			}
		})
	}
}
