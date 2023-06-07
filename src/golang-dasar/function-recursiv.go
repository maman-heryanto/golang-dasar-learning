package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
		// fmt.Println(string(i))
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialLoop(value-1)
	}
}

func main() {
	loop := factorialLoop
	fmt.Println(loop(5))

	loopRecursive :-= factorialRecursive
	fmt.Println(loopRecursive(5))
}
