package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and specify new title")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo by index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.Add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid edit command, please use index:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid index for edit")
			os.Exit(1)
		}
		todos.Edit(index, parts[1])

	case cf.Del != -1:
		todos.Del(cf.Del)

	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)

	case cf.List:
		todos.Print()

	default:
		fmt.Println("Invalid command")
	}
}
