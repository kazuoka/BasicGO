package main

import "fmt"

// ตัวแปร
/*
การประกาศตัวแปร
*/

func main() {
	var myAge int = 42
	myAge1 := 42.5
	var t bool = true
	text := "GeneralKazuoka"
	age1, age2 := 42, 50
	fmt.Println("อายุของฉัน =", myAge)
	fmt.Println("อายุของฉัน =", myAge1)
	fmt.Println(t)
	fmt.Println(text)
	fmt.Println(age1, age2)
}
