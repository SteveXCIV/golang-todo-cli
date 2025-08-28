# Personal Task Manager CLI - 8-Hour Project Plan

## Project Overview
Build a command-line task manager that can create, list, complete, and delete tasks. Tasks will have titles, descriptions, priorities, due dates, and categories. Data will persist to a JSON file.

**Key Go Patterns You'll Learn:**
- Struct composition and methods
- Interface implementation
- Error handling patterns
- JSON marshaling/unmarshaling
- File I/O operations
- Command-line argument parsing
- Time handling

---

## Hour 1: Project Setup and Core Data Structures

**Goals:**
- Set up Go module and project structure
- Define core Task struct and related types
- Create basic constructor functions

**What to focus on:**
- Design your Task struct with appropriate field types
- Consider using Go's time.Time for dates
- Think about how to represent priority (int, string, or custom type?)
- Create a TaskManager struct to hold your collection of tasks

**Hints if stuck:**
- Look up `go mod init` for project initialization
- Research Go struct tags for JSON serialization
- Consider using an enum-like pattern with `iota` for priority levels

---

## Hour 2: File Persistence Layer

**Goals:**
- Implement JSON file reading and writing
- Create methods to load and save task data
- Handle file existence and creation

**What to focus on:**
- Methods on your TaskManager to save/load from JSON
- Graceful handling of missing files
- Proper error propagation

**Hints if stuck:**
- Look into `os.Stat()` for checking file existence
- Research `json.Marshal()` and `json.Unmarshal()` 
- The `os.OpenFile()` function with appropriate flags for file creation

---

## Hour 3: Core Task Operations

**Goals:**
- Implement methods to add, update, and delete tasks
- Create task completion functionality
- Add basic validation

**What to focus on:**
- Methods on TaskManager for CRUD operations
- Use slice operations for managing task collections
- Generate unique IDs for tasks (simple incrementing counter is fine)

**Hints if stuck:**
- Look up slice append patterns and removing elements from slices
- Consider using a simple counter for task IDs stored in your TaskManager
- Research Go's built-in `strings` package for validation

---

## Hour 4: Command-Line Interface Foundation

**Goals:**
- Set up basic CLI argument parsing
- Create a simple command router
- Implement help system

**What to focus on:**
- Parse `os.Args` or use a simple library like `flag`
- Create separate functions for each command (add, list, complete, etc.)
- Basic input validation

**Hints if stuck:**
- The `flag` package is Go's standard library solution for CLI args
- Look into `switch` statements for command routing
- Consider how to handle required vs optional arguments

---

## Hour 5: List and Display Operations

**Goals:**
- Implement task listing with filtering options
- Create formatted output display
- Add sorting capabilities

**What to focus on:**
- Filter tasks by status, priority, or category
- Format output in a readable table-like structure
- Sort by different criteria (due date, priority, etc.)

**Hints if stuck:**
- Look up the `sort` package and `sort.Slice()` function
- Research `fmt.Printf()` for formatted output with padding
- Consider using method receivers to make tasks sortable

---

## Hour 6: Advanced Features and Date Handling

**Goals:**
- Add due date parsing from strings
- Implement overdue task detection
- Add task search functionality

**What to focus on:**
- Parse user-friendly date formats ("tomorrow", "2024-12-25", "+3d")
- Compare dates to find overdue items
- Search tasks by title or description

**Hints if stuck:**
- Look into `time.Parse()` with custom layouts
- Research `time.Now()` and duration arithmetic with `time.Add()`
- The `strings.Contains()` function for text searching

---

## Hour 7: Error Handling and User Experience

**Goals:**
- Implement comprehensive error handling
- Add input validation and user feedback
- Create better CLI help and usage messages

**What to focus on:**
- Wrap errors with context using `fmt.Errorf()`
- Validate task inputs (empty titles, invalid dates)
- Provide clear error messages and usage examples

**Hints if stuck:**
- Look up Go's error wrapping patterns with `%w` verb
- Research creating custom error types
- Consider using `os.Exit()` with different codes for different error types

---

## Hour 8: Polish and Testing

**Goals:**
- Write basic unit tests for core functionality
- Add any missing edge case handling
- Create a simple README with usage examples

**What to focus on:**
- Test your TaskManager methods with different scenarios
- Test JSON serialization/deserialization
- Handle edge cases like empty task lists, invalid IDs

**Hints if stuck:**
- Look into Go's `testing` package and `go test` command
- Research table-driven tests pattern
- Use `testify` package if you want more assertion helpers

---

## Sample Commands and Expected Output

### Adding Tasks
```bash
$ ./taskman add "Buy groceries" --priority high --due tomorrow --category personal
✓ Task added successfully (ID: 1)

$ ./taskman add "Finish Go project" --priority medium --due 2024-12-25
✓ Task added successfully (ID: 2)

$ ./taskman add "Call dentist" --priority low --category health
✓ Task added successfully (ID: 3)
```

### Listing Tasks
```bash
$ ./taskman list
ID | Title              | Priority | Due Date   | Category | Status
---|-------------------|----------|------------|----------|--------
1  | Buy groceries     | HIGH     | 2024-08-21 | personal | pending
2  | Finish Go project | MEDIUM   | 2024-12-25 |          | pending
3  | Call dentist      | LOW      |            | health   | pending

$ ./taskman list --status pending --priority high
ID | Title          | Priority | Due Date   | Category | Status
---|---------------|----------|------------|----------|--------
1  | Buy groceries | HIGH     | 2024-08-21 | personal | pending

$ ./taskman list --category personal
ID | Title          | Priority | Due Date   | Category | Status
---|---------------|----------|------------|----------|--------
1  | Buy groceries | HIGH     | 2024-08-21 | personal | pending
```

### Completing and Deleting Tasks
```bash
$ ./taskman complete 1
✓ Task "Buy groceries" marked as completed

$ ./taskman list
ID | Title              | Priority | Due Date   | Category | Status
---|-------------------|----------|------------|----------|--------
1  | Buy groceries     | HIGH     | 2024-08-21 | personal | completed
2  | Finish Go project | MEDIUM   | 2024-12-25 |          | pending
3  | Call dentist      | LOW      |            | health   | pending

$ ./taskman delete 3
✓ Task "Call dentist" deleted successfully
```

### Searching Tasks
```bash
$ ./taskman search "go"
Found 1 task(s):
ID | Title              | Priority | Due Date   | Category | Status
---|-------------------|----------|------------|----------|--------
2  | Finish Go project | MEDIUM   | 2024-12-25 |          | pending

$ ./taskman search "project"
Found 1 task(s):
ID | Title              | Priority | Due Date   | Category | Status
---|-------------------|----------|------------|----------|--------
2  | Finish Go project | MEDIUM   | 2024-12-25 |          | pending

$ ./taskman search "xyz"
No tasks found matching "xyz"
```

### Handling Overdue Tasks
```bash
$ ./taskman list
ID | Title              | Priority | Due Date   | Category | Status
---|-------------------|----------|------------|----------|--------
2  | Finish Go project | MEDIUM   | 2024-12-25 |          | pending
4  | Review notes      | HIGH     | 2024-08-19 | work     | pending ⚠️ OVERDUE

$ ./taskman list --overdue
ID | Title         | Priority | Due Date   | Category | Status
---|-------------|----------|------------|----------|--------
4  | Review notes | HIGH     | 2024-08-19 | work     | pending ⚠️ OVERDUE
```

### Error Handling Examples
```bash
$ ./taskman complete 999
✗ Error: Task with ID 999 not found

$ ./taskman add ""
✗ Error: Task title cannot be empty

$ ./taskman add "Test task" --due "invalid-date"
✗ Error: Invalid date format. Use YYYY-MM-DD, 'today', 'tomorrow', or '+Xd'

$ ./taskman delete 1
✗ Error: Task with ID 1 not found
```

## Stretch Goals (if you finish early)

- Add task categories/tags
- Implement recurring tasks
- Add colored output using a library like `fatih/color`
- Create task templates
- Add task notes/comments