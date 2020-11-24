package main
//package channelManager

import (
	"container/list"
	"sync"
	"fmt"
)

type channelInfo struct {
	key string
	status int
}

var channelList *list.List
var mutex *sync.Mutex

func Init() {
	channelList = list.New()
	mutex = &sync.Mutex{}
}

func Pushback(key string) *list.Element {
	tmp := new(channelInfo)
	tmp.key = key
	e := channelList.PushBack(tmp)
	return e
}

func Remove(key string) bool {
	defer mutex.Unlock()
	mutex.Lock()
	for e := channelList.Front(); e != nil; e = e.Next() {
		tmp := e.Value.(*channelInfo)
		if tmp.key == key {
			channelList.Remove(e)
			return true
		}
	}
	return false
}

func ShowAll() {
	fmt.Println("start")
	defer mutex.Unlock()
	mutex.Lock()
	for e := channelList.Front(); e != nil; e = e.Next() {
		tmp := e.Value.(*channelInfo)
		fmt.Println(tmp)
	}
	fmt.Println("end")
}

func GetOne() (string, bool) {
	defer mutex.Unlock()
	mutex.Lock()
	for e := channelList.Front(); e != nil; e = e.Next() {
		tmp := e.Value.(*channelInfo)
		if tmp.status == 0 {
			tmp.status = 1
			return tmp.key, true
		}
	}
	return "", false
}

func ReturnOne(key string) bool {
	defer mutex.Unlock()
	mutex.Lock()
	for e := channelList.Front(); e != nil; e = e.Next() {
		tmp := e.Value.(*channelInfo)
		if tmp.key == key {
			tmp.status = 0
			return true
		}
	}
	return false
}

func main() {
	Init()
	//ShowAll()
	Pushback("123")
	//ShowAll()
	Pushback("321")
	//ShowAll()

	//fmt.Println(GetOne())
	//fmt.Println(GetOne())
	//fmt.Println(GetOne())
	//ShowAll()
	//fmt.Println(ReturnOne("123"))
	ShowAll()
	fmt.Println(Remove("123"))
	ShowAll()
}
