# Go TODO CLI

A simple command-line TODO application built while learning Go. This project explores Go's package structure, JSON persistence, and CLI argument parsing.

Read about the development experience in the accompanying blog post: [Giving Golang a Try](https://www.stevexciv.com/blog/giving-golang-a-try/)

## Usage

The CLI supports five main commands: `add`, `list`, `search`, `complete`, and `delete`.

### Adding tasks

```bash
# Add a basic task (defaults to medium priority, due today)
./golang-todo-cli add "Call dentist"

# Add with custom priority and due date
./golang-todo-cli add "Buy groceries" -priority high -due tomorrow

# Add with specific date and category
./golang-todo-cli add "File taxes" -priority high -due 2025-04-15 -category finance

# Add with relative due date
./golang-todo-cli add "Review code" -due +7d
```

### Listing tasks

```bash
# List all tasks
./golang-todo-cli list

# Filter by status
./golang-todo-cli list -status pending
./golang-todo-cli list -status completed

# Filter by priority
./golang-todo-cli list -priority high

# Show only overdue tasks
./golang-todo-cli list -overdue
```

### Searching tasks

```bash
# Search task titles
./golang-todo-cli search "doctor"
```

### Completing tasks

```bash
# Mark task as completed using its ID
./golang-todo-cli complete 1
```

### Deleting tasks

```bash
# Delete a task using its ID
./golang-todo-cli delete 5
```

## Building

```bash
go build
```

## Features

- JSON file persistence (`tasks.db.json`)
- Priority levels (low, medium, high)
- Due date parsing (today, tomorrow, +Xd, yyyy-MM-dd)
- Optional task categories
- Status filtering (pending/completed)
- Task search by title
- Clean tabular output formatting
