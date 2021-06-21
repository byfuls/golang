package main

import (
	"encoding/json"
	"fmt"
)

type Row struct {
	Time      string
	Processor float64
}

func (r *Row) UnmarshalJSON(bs []byte) error {
	arr := []interface{}{}
	json.Unmarshal(bs, &arr)
	// TODO: add error handling here.
	r.Time = arr[0].(string)
	r.Processor = arr[1].(float64)
	return nil
}

func main() {
	rows := []Row{}
	text := `
	[
          ["ooid4", 3],
          ["ooid5", 5]
        ]
	`
	json.Unmarshal([]byte(text), &rows)
	fmt.Println(rows)
	fmt.Println("time: ", rows[0].Time)
	fmt.Println("processor: ", rows[0].Processor)
}
