package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		log.Fatal(err)
	}
	// ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*80))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*80))
	defer cancel()
	req = req.WithContext(ctx)
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Fatal("[response timeout] err: ", err)
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("[read body error] err: ", err)
	}
	log.Println(string(out))
}
