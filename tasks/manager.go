package tasks

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

type ManagerConfig struct {
	Filename string
	Now      func() time.Time
}

type AddTaskRequest struct {
	Title    string
	Priority Priority
	DueDate  DueDate
	Category string
}

type ListTasksRequest struct {
	Status      *Status
	Priority    *Priority
	Category    *string
	OverdueOnly bool
}

type SearchTasksRequest struct {
	Query string
}

type CompleteTaskRequest struct {
	Id int
}

type DeleteTaskRequest struct {
	Id int
}

type Manager struct {
	filename string
	now      func() time.Time
	nextId   int
	tasks    []Task
}

func NewManager() (Manager, error) {
	return NewManagerWithConfig(ManagerConfig{
		Filename: "tasks.json",
		Now:      time.Now,
	})
}

func NewManagerWithConfig(config ManagerConfig) (Manager, error) {
	m := Manager{
		filename: config.Filename,
		now:      config.Now,
		tasks:    []Task{},
	}
	err := m.loadFromFile()
	if err != nil {
		return Manager{}, err
	}
	nextId := 1
	for _, task := range m.tasks {
		if task.Id >= nextId {
			nextId = task.Id + 1
		}
	}
	m.nextId = nextId
	return m, nil
}

func (m *Manager) AddTask(r *AddTaskRequest) error {
	newTask := Task{
		Id:       m.nextId,
		Title:    r.Title,
		Priority: r.Priority,
		DueDate:  r.DueDate,
		Category: r.Category,
		Status:   Pending,
	}
	m.tasks = append(m.tasks, newTask)
	m.nextId++
	return nil
}

func (m *Manager) ListTasks(r *ListTasksRequest) ([]Task, error) {
	tasks := make([]Task, len(m.tasks))
	statusFilter := func(s Status) bool { return true }
	priorityFilter := func(p Priority) bool { return true }
	categoryFilter := func(c string) bool { return true }
	overdueFilter := func(d DueDate) bool { return true }
	for _, task := range m.tasks {
		if !statusFilter(task.Status) {
			continue
		}
		if !priorityFilter(task.Priority) {
			continue
		}
		if !categoryFilter(task.Category) {
			continue
		}
		if !overdueFilter(task.DueDate) {
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (m *Manager) SearchTasks(r *SearchTasksRequest) ([]Task, error) {
	panic("not implemented")
}

func (m *Manager) CompleteTask(r *CompleteTaskRequest) error {
	panic("not implemented")
}

func (m *Manager) DeleteTask(r *DeleteTaskRequest) error {
	panic("not implemented")
}

func (m *Manager) loadFromFile() error {
	if strings.TrimSpace(m.filename) == "" {
		return nil
	}
	file, err := os.ReadFile(m.filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &m.tasks)
}

func (m *Manager) saveToFile() error {
	if strings.TrimSpace(m.filename) == "" {
		return nil
	}
	file, err := json.Marshal(m.tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(m.filename, file, 0644)
}
