package main

import (
  "fmt"
  "time"
  "math/rand"
)
  
func main() {
    rand.Seed(time.Now().UnixNano())
    n := rand.Intn(10) // n will be between 0 and 10
    fmt.Printf("Sleeping %d seconds...\n", n)
    fmt.Printf("duration: %s\n", (time.Duration(n) * time.Second))

    dt := time.Now()
    dtn := time.Now().Add(time.Duration(n) * time.Second)
    fmt.Printf("current time : %s\n", dt.Format("2006-01-02 15:04:05"))
    fmt.Printf("current timen: %s\n", dtn.Format("2006-01-02 15:04:05"))
}
