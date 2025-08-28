package tasks

type ManagerConfig struct{}

type AddTaskRequest struct {
	Title    string
	Priority Priority
	DueDate  DueDate
	Category string
}

type ListTasksRequest struct {
	Status      Status
	Priority    Priority
	Category    string
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

type Manager struct{}

func NewManager() Manager {
	return NewManagerWithConfig(ManagerConfig{})
}

func NewManagerWithConfig(config ManagerConfig) Manager {
	return Manager{}
}

func (m *Manager) AddTask(r *AddTaskRequest) error {
	panic("not implemented")
}

func (m *Manager) ListTasks(r *ListTasksRequest) ([]Task, error) {
	panic("not implemented")
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
