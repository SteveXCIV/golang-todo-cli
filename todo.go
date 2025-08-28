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

type task struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Priority Priority  `json:"priority"`
	DueDate  time.Time `json:"dueDate"`
	Category string    `json:"category"`
}

func main() {
	var testTask = task{
		Id:       1,
		Title:    "Schedule dentist appointment",
		Priority: High,
		DueDate:  time.Now().AddDate(0, 0, 7),
		Category: "Health",
	}
	jsonData, _ := json.Marshal(testTask)
	fmt.Println(string(jsonData))
}
