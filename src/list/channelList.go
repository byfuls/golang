package main

//package channelManager

import (
	"container/list"
	"fmt"
	"sync"
)

type channelInfo struct {
	key    string
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
	fmt.Println("___ showall - start ___")
	defer mutex.Unlock()
	mutex.Lock()
	for e := channelList.Front(); e != nil; e = e.Next() {
		tmp := e.Value.(*channelInfo)
		fmt.Printf("key[%v] status[%v]\n", tmp.key, tmp.status)
	}
	fmt.Println("_____ end _____")
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

	fmt.Println("get one >> ")
	fmt.Println(GetOne())
	ShowAll()
	fmt.Println(ReturnOne("123"))
	ShowAll()
	Remove("123")
	ShowAll()
}
