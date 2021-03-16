package main

import "fmt"

// การดำเนินการทางคณิตศาสตร์

func main() {
	fNumber := 10
	fSecond := 5
	fmt.Println(fNumber + fSecond)
	fmt.Println(fNumber - fSecond)
	fmt.Println(fNumber * fSecond)
	fmt.Println(fNumber / fSecond)

	// ชุดข้อความ
	s1 := "GeneralKazuoka"
	s2 := "Tutorial"
	// Concatenation
	s3 := s1 + s2
	fmt.Println(s3)
	fmt.Println(s3[0:7])
	fmt.Println(s3[7:11])
	fmt.Println(s3[0]) // รหัสแอสกี
	fmt.Println(s3[0:])
}
