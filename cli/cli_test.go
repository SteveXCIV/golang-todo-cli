package cli

import (
	"strings"
	"testing"
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
			if !strings.Contains(err.Error(), "invalid") {
				t.Fatalf("expected 'invlaid' in err, got: %v", err)
			}
		})
	}
}
