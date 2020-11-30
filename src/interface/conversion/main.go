package main

import (
	"fmt"
	"os"
)

func main() {
	result := ReturnData()

	// int(result) wouldn't work at this stage:

	// Type assert, by placing brackets and type after variable name.
	// Note that we need to assign to a new variable.
	myInt, ok := result.(int)
	if !ok {
		fmt.Printf("got data of type %T but wanted int\n", result)
		os.Exit(1)
	}

	// Now we can work with this, should print '10'
	myInt += 5
	fmt.Println(myInt)
}

func ReturnData() interface{} {
	return 5
}
