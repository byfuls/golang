package main

import (
	"encoding/json"
	"fmt"
)

type Row struct {
	Ooid  string
	Score float64
	Text  string
}

func (r *Row) MarshalJSON() ([]byte, error) {
	arr := []interface{}{r.Ooid, r.Score, r.Text}
	return json.Marshal(arr)
}

func (r *Row) UnmarshalJSON(bs []byte) error {
	arr := []interface{}{}
	json.Unmarshal(bs, &arr)
	// TODO: add error handling here.
	r.Ooid = arr[0].(string)
	r.Score = arr[1].(float64)
	r.Text = arr[2].(string)
	return nil
}

func main() {
	rows := []Row{
		{"ooid1", 2.0, "Söme text"},
		{"ooi", 1.3, "Åther text"},
	}
	marshalled, _ := json.Marshal(rows)
	fmt.Println(string(marshalled))

	// Let's go the other way around.
	rows = []Row{}
	text := `
	[
          ["ooid4", 3.1415, "pi"],
          ["ooid5", 2.7182, "euler"]
        ]
	`
	json.Unmarshal([]byte(text), &rows)
	fmt.Println(rows)
	fmt.Println("ooid: ", rows[0].Ooid)
}
