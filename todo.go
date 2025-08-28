package main

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
	if val, ok := stringToPriority[string(b)]; ok {
		*p = val
		return nil
	}
	return fmt.Errorf("unknown priority: %s", string(b))
}

type Status int

const (
	Pending Status = iota
	Completed
)

type DueDate time.Time

func (d DueDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format("2006-01-02"))
}

func (d *DueDate) UnmarshalJSON(b []byte) error {
	if timestamp, err := time.Parse(`"2006-01-02"`, string(b)); err != nil {
		return err
	} else {
		*d = DueDate(timestamp)
		return nil
	}
}

type task struct {
	Id       int      `json:"id"`
	Title    string   `json:"title"`
	Priority Priority `json:"priority"`
	DueDate  DueDate  `json:"dueDate"`
	Category string   `json:"category"`
}

func main() {
	var testTask = task{
		Id:       1,
		Title:    "Schedule dentist appointment",
		Priority: High,
		DueDate:  DueDate(time.Now().AddDate(0, 0, 7)),
		Category: "Health",
	}

	jsonPriority, _ := json.Marshal(testTask.Priority)
	fmt.Println("Marshaled Priority:", string(jsonPriority))

	var parsedPriority Priority
	_ = json.Unmarshal(jsonPriority, &parsedPriority)
	fmt.Println("Parsed Priority:", parsedPriority)

	jsonData, _ := json.Marshal(testTask)
	fmt.Println(string(jsonData))
}
