package cli

import (
	"errors"
	"slices"
	"strings"
	"testing"

	"github.com/stevexciv/golang-todo-cli/tasks"
)

func TestSearchExecuteTableDriven(t *testing.T) {
	m := newMockManager(testTime)
	tests := []struct {
		name       string
		searchCmd  SearchCommand
		mockReturn []tasks.Task
		wantCall   searchCall
		want       string
	}{
		{
			name: "search tasks",
			searchCmd: SearchCommand{
				Query: "tasks",
			},
			mockReturn: []tasks.Task{
				{Id: 1, Title: "Do some tasks"},
			},
			wantCall: searchCall{
				query: "tasks",
			},
			want: "Do some tasks",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.reset()
			m.searchNextOk = tt.mockReturn

			res, err := tt.searchCmd.Execute(&m)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !slices.Contains(m.searchCalls, tt.wantCall) {
				t.Errorf("unexpected search call, wanted %v", tt.wantCall)
			}
			if !strings.Contains(strings.ToLower(res), strings.ToLower(tt.want)) {
				t.Errorf("expected '%v' to contain '%v'", res, tt.want)
			}
		})
	}
}

func TestSearchErrorsTableDriven(t *testing.T) {
	m := newMockManager(testTime)
	tests := []struct {
		name      string
		searchCmd SearchCommand
		mockErr   error
		want      string
	}{
		{
			name: "search tasks error",
			searchCmd: SearchCommand{
				Query: "tasks",
			},
			mockErr: errors.New("search error"),
			want:    "search error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.reset()
			m.searchNextErr = tt.mockErr

			_, err := tt.searchCmd.Execute(&m)

			if err == nil {
				t.Fatalf("expected error: %v", tt.want)
			}
			if !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(tt.want)) {
				t.Errorf("expected '%v' to contain '%v'", err, tt.want)
			}
		})
	}
}
