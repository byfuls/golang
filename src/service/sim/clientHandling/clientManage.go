package clientHandling

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type clientInfo struct {
	updatedTime time.Time
	imsi        string
	addr        net.UDPAddr
}

type clientMgr struct {
	clientList []clientInfo
	mutex      *sync.Mutex
}

func (c *clientMgr) reArrange() {
	now := time.Now()
	for i := 0; i < len(c.clientList); i++ {
		fmt.Printf("[clientMgr/reArrange] (%d)\n", i)
		fmt.Println("[clientMgr/reArrange] time: ", c.clientList[i].updatedTime)

		elapsed := now.Sub(c.clientList[i].updatedTime)
		fmt.Println("[clientMgr/reArrange] elapsed: ", elapsed)
	}
}

func (c *clientMgr) print() {
	for i := 0; i < len(c.clientList); i++ {
		fmt.Printf("[clientMgr/print] (%d)\n", i)
		fmt.Println("[clientMgr/print] time: ", c.clientList[i].updatedTime)
	}
}

func (c *clientMgr) remove(pos int) {

}

func (c *clientMgr) insert(addr net.UDPAddr, imsi string, seq uint32) bool {
	if 0 >= len(imsi) {
		fmt.Println("[insert-err] imsi: ", imsi)
		return false
	}

	client := clientInfo{
		updatedTime: time.Now(),
	}

	defer c.mutex.Unlock()
	c.mutex.Lock()
	c.clientList = append(c.clientList, client)

	c.print()

	return true
}

func (c *clientMgr) init() {
	c.clientList = make([]clientInfo, 0, 1024)
	c.mutex = &sync.Mutex{}
}

/*----------------------------------------------------------------------------*/

var cMgr clientMgr

func Init() {
	cMgr.init()
	go watcher()
}

func watcher() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("[watcher] ... ing ...")
		cMgr.reArrange()
	}
}

func Insert(addr net.UDPAddr, imsi string, seq uint32) bool {
	return cMgr.insert(addr, imsi, seq)
}
