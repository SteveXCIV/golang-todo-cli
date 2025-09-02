package cli

import (
	"time"

	"github.com/stevexciv/golang-todo-cli/tasks"
)

type addCall struct {
	title    string
	priority tasks.Priority
	dueDate  tasks.DueDate
	category string
}

type completeCall struct {
	id int
}

type deleteCall struct {
	id int
}

type listCall struct {
	status      *tasks.Status
	priority    *tasks.Priority
	category    string
	overdueOnly bool
}

type searchCall struct {
	query string
}

type mockManager struct {
	now             time.Time
	addCalls        []addCall
	addNextOk       *tasks.Task
	addNextErr      error
	completeCalls   []completeCall
	completeNextErr error
	deleteCalls     []deleteCall
	deleteNextErr   error
	listCalls       []listCall
	listNextOk      []tasks.Task
	listNextErr     error
	searchCalls     []searchCall
	searchNextOk    []tasks.Task
	searchNextErr   error
}

func (m *mockManager) AddTask(title string, priority tasks.Priority, getDueDate func(time.Time) tasks.DueDate, category string) (*tasks.Task, error) {
	if m.addNextErr != nil {
		return nil, m.addNextErr
	}
	call := addCall{
		title:    title,
		priority: priority,
		dueDate:  getDueDate(m.now),
		category: category,
	}
	m.addCalls = append(m.addCalls, call)
	return m.addNextOk, nil
}

func (m *mockManager) CompleteTask(id int) error {
	if m.completeNextErr != nil {
		return m.completeNextErr
	}
	call := completeCall{
		id: id,
	}
	m.completeCalls = append(m.completeCalls, call)
	return nil
}

func (m *mockManager) DeleteTask(id int) error {
	if m.deleteNextErr != nil {
		return m.deleteNextErr
	}
	call := deleteCall{
		id: id,
	}
	m.deleteCalls = append(m.deleteCalls, call)
	return nil
}

func (m *mockManager) ListTasks(status *tasks.Status, priority *tasks.Priority, category string, overdueOnly bool) ([]tasks.Task, error) {
	if m.listNextErr != nil {
		return []tasks.Task{}, m.listNextErr
	}
	call := listCall{
		status:      status,
		priority:    priority,
		category:    category,
		overdueOnly: overdueOnly,
	}
	m.listCalls = append(m.listCalls, call)
	return m.listNextOk, nil
}

func (m *mockManager) SearchTasks(query string) ([]tasks.Task, error) {
	if m.searchNextErr != nil {
		return []tasks.Task{}, m.searchNextErr
	}
	call := searchCall{
		query: query,
	}
	m.searchCalls = append(m.searchCalls, call)
	return m.searchNextOk, nil
}

func (m *mockManager) resetAdd() {
	m.addCalls = []addCall{}
	m.addNextErr = nil
}

func (m *mockManager) resetComplete() {
	m.completeCalls = []completeCall{}
	m.completeNextErr = nil
}

func (m *mockManager) resetDelete() {
	m.deleteCalls = []deleteCall{}
	m.deleteNextErr = nil
}

func (m *mockManager) resetList() {
	m.listCalls = []listCall{}
	m.listNextErr = nil
	m.listNextOk = []tasks.Task{}
}

func (m *mockManager) resetSearch() {
	m.searchCalls = []searchCall{}
	m.searchNextErr = nil
	m.searchNextOk = []tasks.Task{}
}

func (m *mockManager) reset() {
	m.resetAdd()
	m.resetComplete()
	m.resetDelete()
	m.resetList()
	m.resetSearch()
}

func newMockManager(now time.Time) mockManager {
	return mockManager{
		now: now,
	}
}
