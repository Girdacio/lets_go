package main

import "fmt"

func main() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 30, "Eve": 28}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	delete(myMap2, "Adam")

	var age, ok = myMap2["Eve"]
	if ok {
		fmt.Printf("Eve's age is %d\n", age)
	} else {
		fmt.Println("Eve's age not found")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
}
