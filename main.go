package main

import "errors"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	println(add(1, 2))
	println(sub(1, 2))
	println(mul(1, 2))
	if result, err := div(9, 0); err == nil {
		println(result)
	} else {
		println(err.Error())
	}
}
