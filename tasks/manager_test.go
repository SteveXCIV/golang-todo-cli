package tasks

import (
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	m := Manager{}

	err := m.AddTask(&AddTaskRequest{
		Title:    "Test Task",
		Priority: Medium,
		DueDate:  DueDate(time.Date(2024, 4, 10, 0, 0, 0, 0, time.UTC)),
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(m.tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(m.tasks))
	}
	task := m.tasks[0]
	expectedTask := Task{
		Id:       1,
		Title:    "Test Task",
		Priority: Medium,
		DueDate:  DueDate(time.Date(2024, 4, 10, 0, 0, 0, 0, time.UTC)),
		Status:   Pending,
	}
	if task != expectedTask {
		t.Fatalf("expected %v, got %v", expectedTask, task)
	}
}
