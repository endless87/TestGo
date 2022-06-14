package Data

import (
	"fmt"
	"strings"
)

func GetMessage() {
	fmt.Println("GO 데이터")
}

func LenAndUpper(name string) (int, string) { // return 받을 형 선언
	return len(name), strings.ToUpper(name)
}

func RepeatMe(words ...string) {
	fmt.Println(words)
}

func LenAndUpperNakedReturn(name string) (length int, uppercase string) {
	defer fmt.Println("return success") // return 후에 실행되는 함수
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func SuperAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}
