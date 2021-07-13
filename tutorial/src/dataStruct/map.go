package dataStruct

/*
	HashMap 속도 = O(1)
	SortedMap 속도 = O(log2N)  // BST 방식

	Map + array = O(N)
	Map + BST = log2N //  Sorted Map, Ordered Map
	Map + Hash = O(1)
		Hash
		1. 출력값의 범위가 있다.
		2. 같은 입력이면 같은 출력이 나온다
		3. 다른 입력이면 보통의 경우 다른 출력이 나온다

		sin			modular (*)
		-1~1		0~(N-1)
		정수->실수		정수->정수
			one way function
*/

/* Rolling hash
Hi = (Hi x A + Si) % B
						소수
A : ASCII 0 ~ 255, 256
B : 3571

Hi Range : 0 ~ 3570
*/

func Hash(s string) int {
	h := 0
	A := 256
	B := 3571
	for i := 0; i < len(s); i++ {
		h = (h*A + int(s[i])) % B
	}
	return h
}

type keyValue struct {
	key   string
	value string
}

type Map struct {
	keyArray [3571][]keyValue
}

func (m *Map) Add(key, value string) {
	h := Hash(key)
	m.keyArray[h] = append(m.keyArray[h], keyValue{key, value})
}

func CreateMap() *Map {
	return &Map{}
}

func (m *Map) Get(key string) string {
	h := Hash(key)

	for i := 0; i < len(m.keyArray[h]); i++ {
		if m.keyArray[h][i].key == key {
			return m.keyArray[h][i].value
		}
	}
	return ""
}
