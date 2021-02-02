package main

import "fmt"

func main() {
	var unit string
	fmt.Println("enter (L)bs or (K)g")
	fmt.Scan(&unit)

	var weight float64
	fmt.Println("enter weight")
	fmt.Scan(&weight)

	var converted = weight * 0.45
	var Converted = weight / 0.45


	if unit == "L" {
		fmt.Println("you pet weights", converted, "kilos")
	} else {
		fmt.Println("your pet weights", Converted, "pounds")
	}

}
