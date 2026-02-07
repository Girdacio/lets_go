package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15}
	fmt.Println(myEngine)
	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Printf("Miles left: %d\n", myEngine.milesLeft())

	canMakeIt(myEngine, 50)
	var myElectricEngine electricEngine = electricEngine{3, 10}
	canMakeIt(myElectricEngine, 40)
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Printf("You can make it! You have %d miles left.\n", e.milesLeft())
	} else {
		fmt.Printf("You cannot make it. You only have %d miles left.\n", e.milesLeft())
	}
}
