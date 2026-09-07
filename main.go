package main

func main() {
	todos := Todos{}

	todos.add("Complete assignments.")
	todos.add("Solve DSA Problems.")

	todos.toggle(0)

	todos.print()
}