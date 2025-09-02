package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type Manager interface {
	AddTask(title string, priority Priority, getDueDate func(time.Time) DueDate, category string) (*Task, error)
	ListTasks(status *Status, priority *Priority, category string, overdueOnly bool) ([]Task, error)
	SearchTasks(query string) ([]Task, error)
	CompleteTask(id int) error
	DeleteTask(id int) error
}

type manager struct {
	filename string
	now      func() time.Time
	nextId   int
	tasks    []Task
}

func NewManager() (Manager, error) {
	return newManagerInternal(
		"tasks.json",
		time.Now,
		[]Task{},
	)
}

func newManagerInternal(
	filename string,
	now func() time.Time,
	tasks []Task,
) (Manager, error) {
	if now == nil {
		now = time.Now
	}
	m := &manager{
		filename: filename,
		now:      now,
		tasks:    tasks,
	}
	err := m.loadFromFile()
	if err != nil {
		return nil, err
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

func (m *manager) AddTask(title string, priority Priority, getDueDate func(time.Time) DueDate, category string) (*Task, error) {
	newTask := Task{
		Id:       m.nextId,
		Title:    title,
		Priority: priority,
		DueDate:  getDueDate(m.now()),
		Category: category,
		Status:   Pending,
	}
	m.tasks = append(m.tasks, newTask)
	m.nextId++
	if err := m.saveToFile(); err != nil {
		return nil, err
	}
	return &newTask, nil
}

func (m *manager) ListTasks(status *Status, priority *Priority, category string, overdueOnly bool) ([]Task, error) {
	tasks := make([]Task, 0)
	statusFilter := func(s Status) bool { return true }
	if status != nil {
		statusFilter = func(s Status) bool { return s == *status }
	}
	priorityFilter := func(p Priority) bool { return true }
	if priority != nil {
		priorityFilter = func(p Priority) bool { return p == *priority }
	}
	categoryFilter := func(c string) bool { return true }
	if category != "" {
		categoryFilter = func(c string) bool { return strings.EqualFold(c, category) }
	}
	overdueFilter := func(t *Task) bool { return true }
	if overdueOnly {
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

func (m *manager) SearchTasks(query string) ([]Task, error) {
	tasks := make([]Task, 0)
	queryLower := strings.ToLower(query)
	for _, task := range m.tasks {
		if strings.Contains(strings.ToLower(task.Title), queryLower) {
			tasks = append(tasks, task)
		}
	}
	return tasks, nil
}

func (m *manager) CompleteTask(id int) error {
	for i := range m.tasks {
		if m.tasks[i].Id == id {
			if m.tasks[i].Status == Completed {
				return fmt.Errorf("task %d is already completed", id)
			}
			m.tasks[i].Status = Completed
			return m.saveToFile()
		}
	}
	return fmt.Errorf("task %d not found", id)
}

func (m *manager) DeleteTask(id int) error {
	for i := range m.tasks {
		if m.tasks[i].Id == id {
			m.tasks = slices.Delete(m.tasks, i, i+1)
			return m.saveToFile()
		}
	}
	return fmt.Errorf("task %d not found", id)
}

func (m *manager) loadFromFile() error {
	if strings.TrimSpace(m.filename) == "" {
		return nil
	}
	if _, err := os.Stat(m.filename); os.IsNotExist(err) {
		return nil
	}
	file, err := os.ReadFile(m.filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &m.tasks)
}

func (m *manager) saveToFile() error {
	if strings.TrimSpace(m.filename) == "" {
		return nil
	}
	file, err := json.Marshal(m.tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(m.filename, file, 0644)
}
