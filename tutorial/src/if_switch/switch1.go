package main

import "fmt"

func main() {
	x := 33

	switch { // == switch true {}
	case x > 40:
		fmt.Println("x는 40보다 크다")
	case x < 20:
		fmt.Println("x는 20보다 작다")
	case x > 30: // true, 실행 후 빠져나옴
		fmt.Println("x는 30보다 크다")
	case x == 33: // true, 실행 X
		fmt.Println("x는 33과 같다")
	}
}
