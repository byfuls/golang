package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
	val string
}

func (s *StructA) String() string {
	return "Val: " + s.val
}

type StructB struct {
	val int
}

func (s *StructB) String() string {
	return "StructB: " + strconv.Itoa(s.val)
}

func main() {
	a := &StructA{val: "AAA"}
	fmt.Println(a) // 하동
	Println(a)     // 상동

	b := &StructB{val: 10}
	Println(b)
}

type Printable interface {
	String() string
}

func Println(p Printable) {
	// 인자로 받는 p가 StructA 타입인지 StructB 타입인지 관심이 없다
	// 어떤 것이든 상관없이 interface에 정의된 String() 이라는 관계만 정의 되어 있는지 관심있다
	// 또한 String() 함수가 뭘 하든 관심도 없다, 단순 해당 함수 관계가 있는지만 판단하겠다
	fmt.Println(p.String())
}
