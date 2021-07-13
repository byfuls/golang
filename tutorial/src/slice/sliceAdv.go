package main

import "fmt"

/* slice property
{
	pointer => 시작 주소
	len => 개수
	cap => 최대개수
}
*/

func main() {
	var s []int
	var t []int

	s = make([]int, 3)

	s[0] = 100
	s[1] = 200
	s[2] = 300

	fmt.Println(s, len(s), cap(s))

	t = append(s, 400, 500, 600, 700)
	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))

	fmt.Println("//////////////")

	var u []int
	u = append(t, 500)
	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))
	fmt.Println(u, len(u), cap(u))

	fmt.Println("//////////////")
	/* append 는 대상이 되는 slice의 cap-len 을 값을 확인하고,
	빈 공간이 있는지 판단한다.
	빈 공간이 있다면, 같은 메모리 주소에 처리를 하고,
	빈 공간이 없다면, 새로운 메모리 공간을 할당하여 처리한다
	*/

	u[0] = 9999
	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))
	fmt.Println(u, len(u), cap(u))
}
