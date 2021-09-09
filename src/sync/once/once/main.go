package main

import (
	"log"
	"sync"
)

func main() {
	var once sync.Once    // Once 객체 생성
	var wg sync.WaitGroup // WaitGroup 객체 생성
	for i := 0; i < 3; i++ {
		wg.Add(1)        // 고루틴 등록
		go func(i int) { // 개별 고루틴 생성 및 수행
			defer wg.Done() // 고루틴 종료 시그널
			log.Printf("Goroutine[%v] start\n", i)
			once.Do(func() { // 한번만 실행하게끔 Do() 메서드 수행
				log.Printf("Goroutine[%v] Init Complete", i) // 한번만 실행될 내용
			})
			log.Printf("Goroutine[%v] now starting ...\n", i) // 개별 고루틴 진행
		}(i)
	}
	wg.Wait() // 모든 고루틴 종료까지 Wait
}
