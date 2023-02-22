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

func main() {
	name := "jihye"
	// var name string = "jihye" 위의 방법과 동일
	name = "jisu"

	fmt.Println(name)

	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("jihye")

	fmt.Println(totalLength, upperName)

	repeatMe("jihye", "jisu")
}
