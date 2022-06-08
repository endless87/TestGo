package main

import (
	"TestGo/data"
	"fmt"
)

func main() {

	name := "Soo Kwon"
	message := "Test Go Go"
	fmt.Println(name)
	fmt.Println(message)

	data1 := 1
	data2 := 2

	fmt.Println(data1 + data2)

	data.GetMessage()

	totalLength, upperName := data.LenAndUpper(name)
	fmt.Println(totalLength, upperName)

	data.RepeatMe("one", "two", "three", "four", "five")

	totalLengthP, upperNameP := data.LenAndUpperNakedReturn(name + " " + name)
	fmt.Println(totalLengthP, upperNameP)

	SuperAdd(10, 20, 30, 40, 50)

}

func SuperAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}
