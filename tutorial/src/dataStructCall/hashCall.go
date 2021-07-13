package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	fmt.Println("abcde = ", dataStruct.Hash("abcde"))
	fmt.Println("abcde = ", dataStruct.Hash("abcdf"))
	fmt.Println("abcde = ", dataStruct.Hash("tbcde"))
	fmt.Println("abcde = ", dataStruct.Hash("afdgadfgadgfadsgf"))

	m := dataStruct.CreateMap()
	m.Add("AAA", "0107777777")
	m.Add("BBB", "0108888888")
	m.Add("CDFSFEWFEWFEWF", "0109999999")
	m.Add("CCC", "010123123")

	fmt.Println("AAA = ", m.Get("AAA"))
	fmt.Println("BBB = ", m.Get("BBB"))
	fmt.Println("CCC = ", m.Get("CCC"))
	fmt.Println("DDD = ", m.Get("DDD"))
	fmt.Println("CDFSFEWFEWFEWF = ", m.Get("CDFSFEWFEWFEWF"))
}
