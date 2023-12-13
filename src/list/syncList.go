package main

import (
	"container/list"
	"fmt"
	"sync"
)

type ldata struct {
	key string
	ele any
	ref bool
}

type NodeList struct {
	vlist *list.List
	mtx   *sync.Mutex
}

func NewNodeList() *NodeList {
	return &NodeList{
		vlist: list.New(),
		mtx:   &sync.Mutex{},
	}
}

// 리스트의 마지막 노드에 덧붙혀서 생성(추가)
func (n *NodeList) PushBack(k string, v any) *ldata {
	t := new(ldata)
	t.key = k
	t.ele = v
	t.ref = false
	n.vlist.PushBack(t)
	return t
}

// 리스트의 제일 앞에서 참조하지 않는 노드의 참조만 가져옴
func (n *NodeList) RefFront() (string, any) {
	defer n.mtx.Unlock()
	n.mtx.Lock()
	for e := n.vlist.Front(); e != nil; e = e.Next() {
		t := e.Value.(*ldata)
		if !t.ref {
			t.ref = true
			return t.key, t.ele
		}
	}
	return "", nil
}

// 가져온 노드의 참조를 반납
func (n *NodeList) RefReturn(k string) error {
	defer n.mtx.Unlock()
	n.mtx.Lock()
	return fmt.Errorf("not found the key[%v]", k)
}

// 리스트의 제일 앞에서 참조되지 않는 노드를 삭제 및 가져오기
func (n *NodeList) PopFront() (string, any) {
	defer n.mtx.Unlock()
	n.mtx.Lock()
	for e := n.vlist.Front(); e != nil; e = e.Next() {
		t := e.Value.(*ldata)
		if !t.ref {
			n.vlist.Remove(e)
			return t.key, t.ele
		}
	}
	return "", nil
}

// 리스트의 노드의 수
func (n *NodeList) Len() int {
	return n.vlist.Len()
}

// 리스트의 모든 노드 출력 (디버깅용)
func (n *NodeList) ShowAll() {
	fmt.Println("_________SHOW ALL________")
	defer n.mtx.Unlock()
	n.mtx.Lock()
	i := 0
	for e := n.vlist.Front(); e != nil; e = e.Next() {
		t := e.Value.(*ldata)
		fmt.Printf("[%v] key[%v] ref[%v] ele[%v]\n", i, t.key, t.ref, t.ele)
		i++
	}
	fmt.Println("_________________________")
}

func main() {
	fmt.Println("========== Init & Push Back    ==========")
	l := NewNodeList()
	fmt.Println("len >> ", l.Len())
	l.PushBack("t1", 123)
	fmt.Println("len >> ", l.Len())
	l.PushBack("t2", 2)
	l.PushBack("t3", 3)
	l.PushBack("t4", 4)
	l.ShowAll()
	fmt.Println("=========================================")

	fmt.Println("========== Get Reference Front ==========")
	var (
		k string
		e any
	)
	k, e = l.RefFront()
	fmt.Printf("reference front >> key[%v] ele[%v]\n", k, e)
	l.ShowAll()
	k, e = l.RefFront()
	fmt.Printf("reference front >> key[%v] ele[%v]\n", k, e)
	l.ShowAll()
	fmt.Println("=========================================")

	fmt.Println("========== Get Pop Front       ==========")
	k, e = l.PopFront()
	fmt.Printf("pop front >> key[%v] ele[%v]\n", k, e)
	l.ShowAll()
	fmt.Println("=========================================")
}
