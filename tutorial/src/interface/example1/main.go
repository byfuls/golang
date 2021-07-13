package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
}

func (s *StructA) AAA(x int) int {
	return x * x
}

func (a *StructA) BBB(x int) string {
	return "X= " + strconv.Itoa(x)
}

type StructB struct {
}

func (b *StructB) AAA(x int) int {
	return x * 2
}

func main() {
	var c InterfaceA
	c = &StructA{}

	fmt.Println(c.BBB(3))

	/* Error Case 1 */
	//var d InterfaceA
	//d = &StructB{}		ERR, StructB는 InterfaceA의 BBB()를 구현하고 있지 않기때문에 에러 발생

	/* Error Case 2 */
	//var e InterfaceA
	//e = StructA{}
	// ERR, *StructA 와 StructA는 엄연히 다른 데이터 타입이다.
	// func (a *StructA) AAA(x int) int {} 와				-- 포인터 타입에 대한 메소드
	// func (a StructA) AAA(x int) int {} 는 다른 메소드다		-- 값 타입에 대한 메소드

}
