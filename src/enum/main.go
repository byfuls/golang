package main

import "fmt"

type ByteSize float64

const (
	KB ByteSize = 1 << (10 * (1 + iota))
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println(KB) // 1024
	fmt.Println(MB) // 1.048576e+06
	fmt.Println(GB) // 1.073741824e+09
	fmt.Println(TB) // 1.099511627776e+12
	fmt.Println(PB) // 1.125899906842624e+15
	fmt.Println(EB) // 1.152921504606847e+18
	fmt.Println(ZB) // 1.1805916207174113e+21
	fmt.Println(YB) // 1.2089258196146292e+24
}
