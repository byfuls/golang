package main

import "fmt"

func main() {
	//s := "Hello world"

	//for i := 0; i < len(s); i++ {
	//	fmt.Print(string(s[i]), ", ")
	//}

	/* 한글은 UTF-8에서 1~3 byte를 갖는데
	일반적으로 한글자씩 출력하면 1byte만 출력하기때문에 한글은 깨지게 되어있다.
	rune 타입으로 정의하면 (rune: UTF8 1~3byte 당 하나로 처리) 가능. */
	s := "Hello 월드"
	s2 := []rune(s)
	fmt.Println("len(s2) = ", len(s2))
	for i := 0; i < len(s2); i++ {
		fmt.Print(string(s2[i]), ", ")
	}
}
