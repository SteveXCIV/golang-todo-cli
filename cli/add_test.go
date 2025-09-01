package cli

import (
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
		name     string
		addCmd   AddCommand
		wantCall addCall
		want     string
	}{}
	var manager tasks.Manager = &m

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: revisit this API to see if we can avoid double-pointer
			got, err := tt.addCmd.Execute(&manager)
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
