package dataStruct

import "fmt"

/*	Heap
	- 최대힙 (부모 노드는 자식 노드보다 큰 값이어야 한다) >> 큰 값 찾을때 Max Heap
	- 최소힙 (부모 노드는 자식 노드보다 작은 값이어야 한다)	>> 작은 값 찾을때 Min Heap
	: 최대값, 최소값 찾을대 이점이 있다. or 몇번째 큰값, ...
	: priority queue 우선 순위 큐
	: Heap 트리 Pop을 하면 정렬한 값을 출력 받을 수 있다. (= Heap 정렬)

	- 속도
	: 추가시 , O(Log2(N))
	: 제거시 , O(Log2(N))

	- Heap 정렬 속도
	: push , O(2NLog2M) => O(NLog2M)

	= Array Slice

	N번째 left = 2N + 1
	N번째 right = 2N + 2
	N번째 노드의 부모 = (N-1)/2

	- REF : https://www.youtube.com/watch?v=liJZaku6_KI&list=PLy-g2fnSzUTAaDcLW7hpq0e8Jlt7Zfgd6&index=37
*/

type Heap struct {
	list []int
}

func (h *Heap) Push(v int) {
	h.list = append(h.list, v)

	idx := len(h.list) - 1
	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 {
			break
		}
		if h.list[idx] > h.list[parentIdx] {
			h.list[idx], h.list[parentIdx] = h.list[parentIdx], h.list[idx]
			idx = parentIdx
		} else {
			break
		}
	}
}

func (h *Heap) Print() {
	fmt.Println(h.list)
}

func (h *Heap) Pop() int {
	if len(h.list) == 0 {
		return 0
	}

	top := h.list[0]
	last := h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]

	h.list[0] = last
	idx := 0
	for idx < len(h.list) {
		swapIdx := -1
		leftIdx := idx*2 + 1
		if leftIdx >= len(h.list) {
			break
		}
		if h.list[leftIdx] > h.list[idx] {
			swapIdx = leftIdx
		}

		rightIdx := idx*2 + 2
		if rightIdx < len(h.list) {
			if h.list[rightIdx] > h.list[idx] {
				if swapIdx < 0 || h.list[swapIdx] < h.list[rightIdx] {
					swapIdx = rightIdx
				}
			}
		}

		if swapIdx < 0 {
			break
		}
		h.list[idx], h.list[swapIdx] = h.list[swapIdx], h.list[idx]
		idx = swapIdx
	}
	return top
}
