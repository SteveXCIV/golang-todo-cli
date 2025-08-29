package tasks

import (
	"slices"
	"strings"
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	// TODO: figure out how to make testing more convenient than this
	m := Manager{nextId: 1}

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
	if m.nextId != 2 {
		t.Fatalf("expected nextId to be 2, got %d", m.nextId)
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
	status := Completed
	tasksByStatus, err := m.ListTasks(&ListTasksRequest{
		Status: &status,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(tasksByStatus) != 1 || tasksByStatus[0] != task2 {
		t.Fatalf("expected %v, got %v", task2, tasksByStatus)
	}

	// filter by priority
	priority := High
	tasksByPriority, err := m.ListTasks(&ListTasksRequest{Priority: &priority})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(tasksByPriority) != 1 || tasksByPriority[0] != task3 {
		t.Fatalf("expected %v, got %v", task3, tasksByPriority)
	}

	// filter by category
	category := "Health"
	tasksByCategory, err := m.ListTasks(&ListTasksRequest{Category: &category})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(tasksByCategory) != 1 || tasksByCategory[0] != task1 {
		t.Fatalf("expected %v, got %v", task1, tasksByCategory)
	}

	// filter by category (case insensitive)
	category = "gRoCeRiEs"
	tasksByCategory, err = m.ListTasks(&ListTasksRequest{Category: &category})
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

func TestSearchTasksTableDriven(t *testing.T) {
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
	m := Manager{
		tasks: []Task{
			task1,
			task2,
			task3,
		},
	}
	tests := []struct {
		query    string
		expected []Task
	}{
		{"dentist", []Task{task1}},
		{"tAxEs", []Task{task3}},
		{"t", []Task{task1, task3}},
	}

	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			actual, err := m.SearchTasks(&SearchTasksRequest{Query: tt.query})
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if !slices.Equal(actual, tt.expected) {
				t.Fatalf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestCompleteTask(t *testing.T) {
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
	m := Manager{
		tasks: []Task{
			task1,
			task2,
		},
	}

	// complete task
	err := m.CompleteTask(&CompleteTaskRequest{Id: 1})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if m.tasks[0].Status != Completed {
		t.Fatalf("expected task 1 to be completed, got %v", m.tasks[0].Status)
	}

	// complete already completed task
	err = m.CompleteTask(&CompleteTaskRequest{Id: 2})
	if err == nil {
		t.Fatalf("expected error, got none")
	}
	if !strings.Contains(err.Error(), "already completed") {
		t.Fatalf("expected 'already completed' error, got %v", err)
	}
}

func TestDeleteTask(t *testing.T) {
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
	m := Manager{
		tasks: []Task{
			task1,
			task2,
		},
	}

	// delete task
	err := m.DeleteTask(&DeleteTaskRequest{Id: 1})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if slices.Contains(m.tasks, task1) {
		t.Fatalf("expected task 1 to be deleted, but it still exists")
	}

	// delete unknown task
	err = m.DeleteTask(&DeleteTaskRequest{Id: 999})
	if err == nil {
		t.Fatalf("expected error, got none")
	}
	if !strings.Contains(err.Error(), "not found") {
		t.Fatalf("expected 'not found' error, got %v", err)
	}
}
