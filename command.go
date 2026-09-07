package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cmdflags struct {
	Add    string
	Delete int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *Cmdflags {
	cf := Cmdflags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo; Specify Title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit an existing todo by its index; Specify a new Title. id:new_title")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete an existing todo by its index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle the existing todo's completed field")
	flag.BoolVar(&cf.List, "list", false, "List out all todo's")

	flag.Parse()

	return &cf
}

func (cf *Cmdflags) Execute(todos *Todos){
	switch{
	// for list command
	case cf.List:
		todos.print()

	// for add command
	case cf.Add != "":
		todos.add(cf.Add)

	// for edit command
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2{
			fmt.Println("Error: invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil{
			fmt.Println("Error: invalid index for edits")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	// for toggle command
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	// for delete command
	case cf.Delete != -1:
		todos.delete(cf.Delete)

	default:
		fmt.Println("Invalid Command")
		
	}
}