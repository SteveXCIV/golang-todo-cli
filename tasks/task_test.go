package tasks

import (
	"encoding/json"
	"testing"
)

func TestPriorityJSONRoundTrip(t *testing.T) {
	var tests = []struct {
		priority     Priority
		expectedJSON string
	}{
		{Low, `"LOW"`},
		{Medium, `"MEDIUM"`},
		{High, `"HIGH"`},
	}

	for _, test := range tests {
		marshaled, err := json.Marshal(test.priority)
		if err != nil {
			t.Errorf("failed to marshal priority %v: %v", test.priority, err)
		}

		if string(marshaled) != test.expectedJSON {
			t.Errorf("priority %v: expected %s, got %s", test.priority, test.expectedJSON, string(marshaled))
		}

		var unmarshaled Priority
		err = json.Unmarshal(marshaled, &unmarshaled)

		if err != nil {
			t.Errorf("failed to unmarshal priority %v: %v", test.priority, err)
		}

		if test.priority != unmarshaled {
			t.Errorf("priority %v != unmarshaled %v", test.priority, unmarshaled)
		}
	}
}

func TestStatusJSONRoundTrip(t *testing.T) {
	var tests = []struct {
		status       Status
		expectedJSON string
	}{
		{Pending, `"pending"`},
		{Completed, `"completed"`},
	}

	for _, test := range tests {
		marshaled, err := json.Marshal(test.status)
		if err != nil {
			t.Errorf("failed to marshal status %v: %v", test.status, err)
		}

		if string(marshaled) != test.expectedJSON {
			t.Errorf("status %v: expected %s, got %s", test.status, test.expectedJSON, string(marshaled))
		}

		var unmarshaled Status
		err = json.Unmarshal(marshaled, &unmarshaled)

		if err != nil {
			t.Errorf("failed to unmarshal status %v: %v", test.status, err)
		}

		if test.status != unmarshaled {
			t.Errorf("status %v != unmarshaled %v", test.status, unmarshaled)
		}
	}
}
