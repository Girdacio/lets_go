package main

import "fmt"

func main() {
	printMessage("Hello from Tutorial 2!")
}

func printMessage(message string) {
	fmt.Println(message)

	result, err := intDivisionTreatment(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(err.Error())
	} else {
		fmt.Println("Result with treatment:", result)
	}

	var simpleResult = intDivisionSimple(10, 2)
	fmt.Println("Result without treatment:", simpleResult)
}

func intDivisionTreatment(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

func intDivisionSimple(a, b int) int {
	return a / b
}
