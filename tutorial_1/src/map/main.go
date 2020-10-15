package main

import "fmt"

type Key struct {
	v int
}

type Value struct {
	v int
}

func main() {
	/* map[키(타입))]값(타입) */
	//var m1 map[int]string
	//var m2 map[Key]Value
	//var m3 map[Key]*Value
	var m map[string]string
	m = make(map[string]string) // 초기화 반드시 필요
	m["bcd"] = "bbb"
	fmt.Println(m["bcd"])

	m1 := make(map[int]string)
	m1[53] = "ddd"
	fmt.Println(m1[53])

	fmt.Println("m1[55] = ", m1[55]) // 초기값 "" 출력

	m2 := make(map[int]int)
	m2[4] = 4
	fmt.Println("m2[10] = ", m2[10]) // 초기값 0 출력

	m2_val1, ok1 := m2[4]
	m2_val2, ok2 := m2[10]
	fmt.Println(m2_val1, ok1)
	fmt.Println(m2_val2, ok2)

	m3 := make(map[int]bool)
	m3[4] = true
	fmt.Println(m3[6], m3[4])

	/*------------------------------------*/
	/* delete(map 이름, 해당 key값) */
	delete(m2, 4)
	m2_val1, ok1 = m2[4]
	fmt.Println(m2_val1, ok1)

	/*------------------------------------*/
	m2[2] = 2
	m2[10] = 10
	m2[11] = 11
	m2[21] = 21
	m2[51] = 51

	for key, value := range m2 {
		fmt.Println(key, " ", value)
	}
}
