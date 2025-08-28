package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/stevexciv/golang-todo-cli/tasks"
)

func main() {
	var testTask = tasks.Task{
		Id:       1,
		Title:    "Schedule dentist appointment",
		Priority: tasks.High,
		DueDate:  tasks.DueDate(time.Now().AddDate(0, 0, 7)),
		Category: "Health",
		Status:   tasks.Pending,
	}

	jsonStatus, _ := json.Marshal(tasks.Completed)
	fmt.Println("Marshaled Status:", string(jsonStatus))

	var parsedStatus tasks.Status
	_ = json.Unmarshal(jsonStatus, &parsedStatus)
	fmt.Println("Parsed Status:", parsedStatus)

	jsonData, _ := json.Marshal(testTask)
	fmt.Println(string(jsonData))
}
