package tasks

import (
	"encoding/json"
	"testing"
)

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
