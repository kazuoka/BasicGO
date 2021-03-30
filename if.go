package main

import "fmt"

/*
func main() {
	fmt.Print("Input Your Number: ")
	var input float64
	fmt.Scanf("%f", &input)
	codition := input > 2
	if codition {
		fmt.Print("Worked")
	} else {
		fmt.Print("Not Worked")
	}
}

*/
func main() {
	// && 2 เงื่อนไขเป็นจริงทั้งคู่
	// || เงื่อนไขใดเป็นจริง 1 เงื่อนไข

	/* if 6 > 3 && 8 > 5 {
		fmt.Print("Worked")
	} else {
		fmt.Print("Not Worked")
	} */

	if 6 > 3 || 8 < 5 {
		fmt.Print("Worked")
	} else {
		fmt.Print("Not Worked")
	}
}
