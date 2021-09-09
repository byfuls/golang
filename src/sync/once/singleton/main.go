package main

import (
	"log"
	"sync"
)

var instance *goroutineCounter
var once sync.Once

type goroutineCounter struct {
	count int
}

func GetInstance() *goroutineCounter {
	once.Do(func() {
		log.Println("GetInstance Init Complete")
		instance = &goroutineCounter{count: 0}
	})
	return instance
}

func (g *goroutineCounter) add() {
	g.count++
}

func (g *goroutineCounter) get() int {
	return g.count
}

func main() {
	var wg sync.WaitGroup // WaitGroup 객체 생성
	for i := 0; i < 3; i++ {
		wg.Add(1)        // 고루틴 등록
		go func(i int) { // 개별 고루틴 생성 및 수행
			defer wg.Done() // 고루틴 종료 시그널
			log.Printf("Goroutine[%v] start\n", i)
			GetInstance().add()
			log.Printf("Goroutine[%v] Init Complete(count:%v)\n", i, GetInstance().get()) // 한번만 실행될 내용

			log.Printf("Goroutine[%v] now starting ...\n", i) // 개별 고루틴 진행
		}(i)
	}
	wg.Wait() // 모든 고루틴 종료까지 Wait
}
