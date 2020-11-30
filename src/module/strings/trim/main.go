package main

import (
	"fmt"
	"strings"
)

func main() {
	tmp := [4]byte{0x31, 0x32, 0x33, 0x00}
	rcv := string(tmp[:])
	ttt := strings.Trim(rcv, "\x00")

	fmt.Printf("[ori] tmp: (%d)[%s]\n", len(tmp), tmp)
	fmt.Printf("[cov] rcv: (%d)[%s]\n", len(rcv), rcv)
	fmt.Printf("[cov] ttt: (%d)[%s]\n", len(ttt), ttt)
}
