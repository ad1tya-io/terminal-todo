package main

import "fmt"

func main() {
	todos := Todos{}

	//adding the items:
	todos.add("Complete assignments.")
	todos.add("Solve DSA Problems.")

	fmt.Printf("%+v\n\n", todos)

	//deleting the item:
	todos.delete(0)
	fmt.Printf("%+v\n\n", todos)
}