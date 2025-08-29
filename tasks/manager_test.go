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

func TestListTasksWithFilters(t *testing.T) {
	task1 := Task{
		Id:       1,
		Title:    "Call dentist",
		Priority: Medium,
		DueDate:  DueDate(time.Date(2024, 4, 10, 0, 0, 0, 0, time.UTC)),
		Category: "Health",
		Status:   Pending,
	}
	task2 := Task{
		Id:       2,
		Title:    "Buy milk",
		Priority: Low,
		DueDate:  DueDate(time.Date(2024, 4, 10, 0, 0, 0, 0, time.UTC)),
		Category: "Groceries",
		Status:   Completed,
	}
	task3 := Task{
		Id:       3,
		Title:    "File taxes",
		Priority: High,
		DueDate:  DueDate(time.Date(2024, 4, 12, 0, 0, 0, 0, time.UTC)),
		Category: "Finance",
		Status:   Pending,
	}
	nowFunc := func() time.Time { return time.Date(2024, 4, 11, 0, 0, 0, 0, time.UTC) }
	m := Manager{
		now: nowFunc,
		tasks: []Task{
			task1,
			task2,
			task3,
		},
	}

	// filter by status
	tasksByStatus, err := m.ListTasks(&ListTasksRequest{
		Status: Completed,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(tasksByStatus) != 1 || tasksByStatus[0] != task2 {
		t.Fatalf("expected %v, got %v", task2, tasksByStatus)
	}

	// filter by priority
	tasksByPriority, err := m.ListTasks(&ListTasksRequest{Priority: High})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(tasksByPriority) != 1 || tasksByPriority[0] != task3 {
		t.Fatalf("expected %v, got %v", task3, tasksByPriority)
	}

	// filter by category
	tasksByCategory, err := m.ListTasks(&ListTasksRequest{Category: "Health"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(tasksByCategory) != 1 || tasksByCategory[0] != task1 {
		t.Fatalf("expected %v, got %v", task1, tasksByCategory)
	}

	// filter by category (case insensitive)
	tasksByCategory, err = m.ListTasks(&ListTasksRequest{Category: "gRoCeRiEs"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(tasksByCategory) != 1 || tasksByCategory[0] != task2 {
		t.Fatalf("expected %v, got %v", task2, tasksByCategory)
	}

	// filter by overdue
	tasksOverdue, err := m.ListTasks(&ListTasksRequest{OverdueOnly: true})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(tasksOverdue) != 1 || tasksOverdue[0] != task1 {
		t.Fatalf("expected %v, got %v", task1, tasksOverdue)
	}
}
