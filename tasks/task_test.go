package tasks

import (
	"encoding/json"
	"testing"
	"time"
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

func TestDueDateJSONRoundTrip(t *testing.T) {
	date := DueDate(time.Date(2024, 4, 10, 0, 0, 0, 0, time.UTC))
	expectedJSON := `"2024-04-10"`

	marshaled, err := json.Marshal(date)
	if err != nil {
		t.Errorf("failed to marshal due date: %v", err)
	}

	if string(marshaled) != expectedJSON {
		t.Errorf("due date: expected %s, got %s", expectedJSON, string(marshaled))
	}

	var unmarshaled DueDate
	err = json.Unmarshal(marshaled, &unmarshaled)
	if err != nil {
		t.Errorf("failed to unmarshal due date: %v", err)
	}

	if date != unmarshaled {
		t.Errorf("due date %v != unmarshaled %v", date, unmarshaled)
	}
}

func TestTaskJSONRoundTrip(t *testing.T) {
	task := Task{
		Id:       1,
		Title:    "Make vet appointment",
		Priority: High,
		DueDate:  DueDate(time.Date(2025, 10, 10, 0, 0, 0, 0, time.UTC)),
		Category: "pet stuff",
		Status:   Pending,
	}

	marshaled, err := json.Marshal(task)
	if err != nil {
		t.Errorf("failed to marshal task: %v", err)
	}

	var unmarshaled Task
	err = json.Unmarshal(marshaled, &unmarshaled)
	if err != nil {
		t.Errorf("failed to unmarshal task: %v", err)
	}

	if task != unmarshaled {
		t.Errorf("task %v != unmarshaled %v", task, unmarshaled)
	}
}
