package main

import (
	"fmt"
	"os"

	"github.com/stevexciv/golang-todo-cli/cli"
	"github.com/stevexciv/golang-todo-cli/tasks"
)

func main() {
	args := os.Args[1:]
	cmd, err := cli.Parse(&args)
	if err != nil {
		fmt.Println("Error parsing command: ", err)
		os.Exit(1)
	}

	m, err := tasks.NewManager()
	if err != nil {
		fmt.Println("Error creating task manager: ", err)
		os.Exit(1)
	}

	res, err := cmd.Execute(m)
	if err != nil {
		fmt.Println("Error executing command: ", err)
		os.Exit(1)
	}

	fmt.Println(res)
}
