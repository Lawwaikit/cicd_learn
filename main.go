package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go-ci-demo <add|sub|mul|div> num1 num2")
		os.Exit(1)
	}

	op := os.Args[1]
	a, _ := strconv.Atoi(os.Args[2])
	b, _ := strconv.Atoi(os.Args[3])

	var result int
	switch op {
	case "add":
		result = Add(a, b)
	case "sub":
		result = Sub(a, b)
	case "mul":
		result = Mul(a, b)
	case "div":
		result = Div(a, b)
	default:
		fmt.Println("Unknown operation:", op)
		os.Exit(1)
	}

	fmt.Printf("Result: %d\n", result)
}
