package main

import "fmt"

/* interface
= A 객체와 B 객체간의 상호관계에 대해 따로 정의한 것을 interface라 한다.
= 객체는 상태와 기능인데 기능 부분만 따로 정의한 것을 interface
*/

/* Object
= 상태 + 기능
	기능 1. 공개기능	-> 관계 -> interface
		2. 내부기능
*/

type SpoonOfJam interface {
	String() string
}

type Jam interface { // 관계만 정의한다
	//GetOneSpoon() *SpoonOfStrawberryJam
	GetOneSpoon() SpoonOfJam
}

type Bread struct {
	val string
}

type AppleJam struct {
}

func (a *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

type StrawberryJam struct {
}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

type OrangeJam struct {
}

func (o *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

type SpoonOfStrawberryJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry"
}

type SpoonOfOrangeJam struct {
}

func (s *SpoonOfOrangeJam) String() string {
	return "+ Orange"
}

type SpoonOfAppleJam struct {
}

func (s *SpoonOfAppleJam) String() string {
	return "+ Apple"
}

//func (b *Bread) PutJam(jam *StrawberryJam) {
func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread " + b.val
}

func main() {
	bread := &Bread{}
	//jam := &StrawberryJam{}
	//jam := &OrangeJam{}
	jam := &AppleJam{}
	bread.PutJam(jam)

	fmt.Println(bread)
}
