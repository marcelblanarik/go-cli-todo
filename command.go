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
	Delete int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and specify new title")
	flag.IntVar(&cf.Delete, "del", -1, "Delete a todo by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo by index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func Help() {
	fmt.Println("Usage: todo [options]")
	fmt.Println("Options:")
	fmt.Println("  -add string")
	fmt.Println("        Add a new todo")
	fmt.Println("  -del int")
	fmt.Println("        Delete a todo by index")
	fmt.Println("  -edit string")
	fmt.Println("        Edit a todo by index and specify new title")
	fmt.Println("  -list")
	fmt.Println("        List all todos")
	fmt.Println("  -toggle int")
	fmt.Println("        Toggle a todo by index")
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

	case cf.Delete != -1:
		todos.Delete(cf.Delete)

	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)

	case cf.List:
		todos.Print()

	default:
		Help()
	}
}
