package main

/* OOP
object = 상태 + 기능(Method(=func))
*/

type Bread sturct {
	val string
}

type Jam struct {

}

func (b *Bread) PutJam(jam *Jam) {
	b.val += jam.GetVal()
}


func (j *Jam) Getval() string {
	return "jam"
}

func main() {

}