package tasks

import (
	"encoding/json"
	"fmt"
	"time"
)

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

var priorityToString = map[Priority]string{
	Low:    "LOW",
	Medium: "MEDIUM",
	High:   "HIGH",
}

var stringToPriority = map[string]Priority{
	"LOW":    Low,
	"MEDIUM": Medium,
	"HIGH":   High,
}

func (p *Priority) String() string {
	return priorityToString[*p]
}

func (p Priority) MarshalJSON() ([]byte, error) {
	if str, ok := priorityToString[p]; ok {
		return json.Marshal(str)
	}
	return nil, fmt.Errorf("unknown priority: %d", p)
}

func (p *Priority) UnmarshalJSON(b []byte) error {
	var k string
	if err := json.Unmarshal(b, &k); err != nil {
		return fmt.Errorf("priority should be a string, got %s", string(b))
	}
	if val, ok := stringToPriority[k]; ok {
		*p = val
		return nil
	}
	return fmt.Errorf("unknown priority: %s", k)
}

type Status int

const (
	Pending Status = iota
	Completed
)

var statusToString = map[Status]string{
	Pending:   "pending",
	Completed: "completed",
}

var stringToStatus = map[string]Status{
	"pending":   Pending,
	"completed": Completed,
}

func (s *Status) String() string {
	if str, ok := statusToString[*s]; ok {
		return str
	}
	return "UNKNOWN"
}

func (s Status) MarshalJSON() ([]byte, error) {
	if str, ok := statusToString[s]; ok {
		return json.Marshal(str)
	}
	return nil, fmt.Errorf("unknown status: %d", s)
}

func (s *Status) UnmarshalJSON(b []byte) error {
	var k string
	if err := json.Unmarshal(b, &k); err != nil {
		return fmt.Errorf("status should be a string, got %s", string(b))
	}
	if val, ok := stringToStatus[k]; ok {
		*s = val
		return nil
	}
	return fmt.Errorf("unknown status: %s", k)
}

type DueDate time.Time

func (d DueDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format("2006-01-02"))
}

func (d *DueDate) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return fmt.Errorf("dueDate should be a string, got %s", string(b))
	}
	if timestamp, err := time.Parse("2006-01-02", v); err != nil {
		return err
	} else {
		*d = DueDate(timestamp)
		return nil
	}
}

type Task struct {
	Id       int      `json:"id"`
	Title    string   `json:"title"`
	Priority Priority `json:"priority"`
	DueDate  DueDate  `json:"dueDate"`
	Category string   `json:"category"`
	Status   Status   `json:"status"`
}
