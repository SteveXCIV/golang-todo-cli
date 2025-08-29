package tasks

import (
	"slices"
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

func TestListTasksNoFilter(t *testing.T) {
	m := Manager{
		tasks: []Task{
			{
				Id:       1,
				Title:    "Test Task",
				Priority: Medium,
				DueDate:  DueDate(time.Date(2024, 4, 10, 0, 0, 0, 0, time.UTC)),
				Status:   Pending,
			},
			{
				Id:       2,
				Title:    "Test Task 2",
				Priority: Low,
				DueDate:  DueDate(time.Date(2024, 4, 11, 0, 0, 0, 0, time.UTC)),
				Status:   Pending,
			},
		},
	}

	tasks, err := m.ListTasks(&ListTasksRequest{})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}
	expectedTasks := []Task{
		{
			Id:       1,
			Title:    "Test Task",
			Priority: Medium,
			DueDate:  DueDate(time.Date(2024, 4, 10, 0, 0, 0, 0, time.UTC)),
			Status:   Pending,
		},
		{
			Id:       2,
			Title:    "Test Task 2",
			Priority: Low,
			DueDate:  DueDate(time.Date(2024, 4, 11, 0, 0, 0, 0, time.UTC)),
			Status:   Pending,
		},
	}
	if !slices.Equal(tasks, expectedTasks) {
		t.Fatalf("expected %v, got %v", expectedTasks, tasks)
	}
}
