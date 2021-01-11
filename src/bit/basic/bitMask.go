package main

import (
	"fmt"
)

type Bits uint8

const (
	Val_1 Bits = 1 << iota
	Val_2
	Val_4
	Val_8
	Val_16
	Val_32
)

func Set(b, flag Bits) Bits    { return b | flag }
func Clear(b, flag Bits) Bits  { return b &^ flag }
func Toggle(b, flag Bits) Bits { return b ^ flag }
func Has(b, flag Bits) bool    { return b&flag != 0 }

func main() {
	fmt.Printf("===========\n")
	fmt.Printf("Val_1: %x\n", Val_1)
	fmt.Printf("Val_2: %x\n", Val_2)
	fmt.Printf("Val_4: %x\n", Val_4)
	fmt.Printf("===========\n")

	var b Bits
	b = Set(b, Val_1)
	b = Toggle(b, Val_4)

	fmt.Printf("_________________\n")
	for i, flag := range []Bits{Val_1, Val_2, Val_4} {
		fmt.Println(i, Has(b, flag))
	}

	b = Toggle(b, Val_4)
	fmt.Printf("_________________\n")
	for i, flag := range []Bits{Val_1, Val_2, Val_4} {
		fmt.Println(i, Has(b, flag))
	}
}
