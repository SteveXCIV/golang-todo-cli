package tasks

import (
	"encoding/json"
	"fmt"
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
	tasks := make([]Task, 0)
	statusFilter := func(s Status) bool { return true }
	if r.Status != nil {
		statusFilter = func(s Status) bool { return s == *r.Status }
	}
	priorityFilter := func(p Priority) bool { return true }
	if r.Priority != nil {
		priorityFilter = func(p Priority) bool { return p == *r.Priority }
	}
	categoryFilter := func(c string) bool { return true }
	if r.Category != nil {
		categoryFilter = func(c string) bool { return strings.EqualFold(c, *r.Category) }
	}
	overdueFilter := func(t *Task) bool { return true }
	if r.OverdueOnly {
		overdueFilter = func(t *Task) bool {
			if t.Status == Completed {
				return false
			}
			return time.Time(t.DueDate).Before(m.now())
		}
	}
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
		if !overdueFilter(&task) {
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (m *Manager) SearchTasks(r *SearchTasksRequest) ([]Task, error) {
	tasks := make([]Task, 0)
	query := strings.ToLower(r.Query)
	for _, task := range m.tasks {
		if strings.Contains(strings.ToLower(task.Title), query) {
			tasks = append(tasks, task)
		}
	}
	return tasks, nil
}

func (m *Manager) CompleteTask(r *CompleteTaskRequest) error {
	for i := range m.tasks {
		if m.tasks[i].Id == r.Id {
			if m.tasks[i].Status == Completed {
				return fmt.Errorf("task %d is already completed", r.Id)
			}
			m.tasks[i].Status = Completed
			return nil
		}
	}
	return fmt.Errorf("task %d not found", r.Id)
}

func (m *Manager) DeleteTask(r *DeleteTaskRequest) error {
	for i := range m.tasks {
		if m.tasks[i].Id == r.Id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task %d not found", r.Id)
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
