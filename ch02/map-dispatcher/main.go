package main

import "fmt"

type calculateFunc func(int, int) int

var (
	operations = map[string]calculateFunc{
		"+": add,
		"-": sub,
		"*": mult,
		"/": div,
	}
)

func main() {
	fmt.Println(calculate(2, 3, "+"))
	fmt.Println(calculateWithMap(2, 3, "-"))
}

func calculateWithMap(a, b int, op string) int {
	if f, ok := operations[op]; ok {
		return f(a, b)
	}
	panic("unsupported operation")
}

func calculate(a, b int, op string) int {
	switch op {
	case "+":
		return add(a, b)
	case "-":
		return sub(a, b)
	case "*":
		return mult(a, b)
	case "/":
		return div(a, b)
	default:
		panic("unsupported operation")
	}
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mult(a, b int) int {
	return a * b
}
func div(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}
