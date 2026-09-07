package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct{
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time
}

type Todos []Todo 

// add method:
func(todos *Todos) add(title string){
	todo := Todo{
		Title : title,
		Completed: false,
		CompletedAt: nil,
		CreatedAt: time.Now(),
	}

	*todos = append(*todos, todo)
}
func(todos *Todos) validateIndex(index int) error{
	if index < 0 || index >= len(*todos){
		err := errors.New("Invalid Index ")
		fmt.Println(err)
		return err
	}
	return nil
}

// delete method
func(todos *Todos) delete(index int) error{
	t:= *todos

	if err := t.validateIndex(index); err != nil{
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil

}
// toggle method
func(todos *Todos) toggle(index int) error{
	t:= *todos

	if err:= t.validateIndex(index); err != nil{
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted{
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}
	t[index].Completed = !isCompleted
	return nil
}
// edit method
func(todos *Todos) edit(index int, title string) error{
	t:= *todos

	if err:= t.validateIndex(index); err != nil{
		return err
	}

	t[index].Title = title

	return nil
}
// print the todo list
func(todos *Todos) print(){
	table := table.New(os.Stdout)

	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "CretedAt", "CompletedAt")

	for index, t := range *todos{
		completed := "❌"
		completedAt := ""
		if t.Completed{
			completed = "✔️"
			if t.CompletedAt != nil{
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)

	}

	table.Render()
}
