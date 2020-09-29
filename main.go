package main

import "fmt"

func main()  {
	var Base uint64
	var altura uint64

	fmt.Scanln(&Base)
	fmt.Scanln(&altura)
	area := (Base*altura)/2
	fmt.Println(area)
}