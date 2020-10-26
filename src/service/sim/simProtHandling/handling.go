package simProtHandling

import (
	"fmt"
	"service/sim/clientHandling"
)

func AS07(pdata *parsedData) {
	fmt.Println("[AS07] start")
	fmt.Println(pdata)
	if ret := clientHandling.Insert(pdata.Addr, pdata.Head.Imsi, pdata.Head.Seq); ret == false {
		fmt.Println("[AS07] insert error")
	}
}
