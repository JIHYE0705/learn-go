package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

//func lenAndUpper(name string) (int, string) {
//	// multiple value 를 반환하는 function
//	return len(name), strings.ToUpper(name)
//}

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	return true
}

func canIDrinkSwitch(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	name := "jihye"
	// var name string = "jihye" 위의 방법과 동일
	name = "jisu"

	fmt.Println(name)

	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("jihye")

	fmt.Println(totalLength, upperName)

	repeatMe("jihye", "jisu")

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	fmt.Println(canIDrink(16))

	fmt.Println(canIDrinkSwitch(18))
}
