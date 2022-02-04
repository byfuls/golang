package main

import (
	"log"
	"reflect"
	"sync"
)

func isLocked(m *sync.Mutex) bool {
	state := reflect.ValueOf(m).Elem().FieldByName("state")
	log.Println(state)
	return state.Int()&1 == 1
}

func main() {
	var mutex = &sync.Mutex{}

	// mutex.Lock()
	log.Println("isLocked? ", isLocked(mutex))
	log.Println("hi")
	mutex.Unlock()

}
