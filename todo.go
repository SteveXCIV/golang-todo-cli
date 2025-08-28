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

type Status int

const (
	Pending Status = iota
	Completed
)

type DueDate time.Time

// JSON Marshal/Unmarshal
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
	jsonDueDate, _ := json.Marshal(testTask.DueDate)
	fmt.Println("Marshaled DueDate:", string(jsonDueDate))
	var parsedDueDate DueDate
	_ = json.Unmarshal(jsonDueDate, &parsedDueDate)
	fmt.Println("Parsed DueDate:", time.Time(parsedDueDate).Format("2006-01-02"))

	jsonData, _ := json.Marshal(testTask)
	fmt.Println(string(jsonData))
}
