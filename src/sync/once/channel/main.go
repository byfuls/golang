package main

import (
	"log"
	"sync"
)

func main() {
	done := make(chan struct{})
	go func() { // 초기화 진행용 고루틴! 한번만 수행
		defer close(done)            // 고루틴 종료 전, 채널로 데이터(=시그널) 전송
		log.Println("Init Complete") // 초기화 진행
	}()

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)        // WaitGroup에 고루틴 등록
		go func(i int) { // 분산 처리 고루틴, 초기화 진행용 고루틴 이후에 진행
			defer wg.Done() // 고루틴 종료 전, 고루틴 종료 시그널 발생
			log.Printf("Goroutine[%v] start\n", i)
			<-done                                            // 채널로부터 데이터(=시그널) 수신
			log.Printf("Goroutine[%v] now starting ...\n", i) // 개별 고루틴 진행
		}(i)
	}
	wg.Wait() // 모든 고루틴 종료될 때까지 대기(blocking)
}
