//package main
package channelManager

import (
	"fmt"
)

type tmpStruct struct {
	a string
	b int
}

var chMgr map[string]interface{}

/*******************************************************************************/

func Del(key string) bool {
	if 0 >= len(key) {
		return false
	}

	delete(chMgr, key)
	return true
}

func Get(key string) (interface{}, bool) {
	if 0 >= len(key) {
		return nil, false
	}
	if 0 >= len(chMgr) {
		return nil, false
	}

	return chMgr[key], true
}

func Put(key string, value interface{}) bool {
	if 0 >= len(key) {
		return false
	}
	fmt.Println("[manager] put: ", value)
	chMgr[key] = value
	return true
}

func Init() {
	chMgr = make(map[string]interface{})

	fmt.Println(chMgr)
}

//func main() {
//	Init()
//
//	chMgr["key"] = tmpStruct{
//		a: "a",
//		b: 3,
//	}
//
//	fmt.Println(chMgr)
//
//	get := chMgr["key"].(tmpStruct)
//	fmt.Println(get)
//	fmt.Println(get.a)
//}
