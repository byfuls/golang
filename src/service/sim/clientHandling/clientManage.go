package clientHandling

import (
	"container/list"
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
	clientList *list.List
	mutex      *sync.Mutex

	imsi map[string]*clientInfo
}

func (c *clientMgr) reArrange() {
	now := time.Now()
	defer c.mutex.Unlock()
	c.mutex.Lock()
	for e := c.clientList.Front(); e != nil; e = e.Next() {
		v := e.Value.(*clientInfo)
		fmt.Println("[clientMgr/reArrange] time: ", v.updatedTime)

		elapsed := now.Sub(v.updatedTime)
		fmt.Println("[clientMgr/reArrange] elapsed: ", elapsed)

		if int(elapsed.Seconds()) > 10 {
			fmt.Println("[clientMgr/reArrange] TIMEOUT")
			c.remove(e, v.imsi)
		}
	}
}

func (c *clientMgr) print() {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	for e := c.clientList.Front(); e != nil; e = e.Next() {
		v := e.Value.(*clientInfo)
		fmt.Println("[clientMgr/reArrange] time: ", v.updatedTime)
	}

	for key, element := range c.imsi {
		fmt.Println("Key: ", key, " => ", "Element: ", element)
		fmt.Println("updated time: ", element.updatedTime)
		fmt.Println("updated imsi: ", element.imsi)
	}
}

func (c *clientMgr) remove(element *list.Element, imsi string) {
	delete(c.imsi, imsi)
	c.clientList.Remove(element)
}

func (c *clientMgr) insert(addr net.UDPAddr, imsi string, seq uint32) bool {
	if 0 >= len(imsi) {
		fmt.Println("[insert-err] imsi: ", imsi)
		return false
	}

	fmt.Println("[insert] imsi: ", imsi)
	fmt.Println("[insert] seq: ", seq)

	client := new(clientInfo)
	client.updatedTime = time.Now()
	client.imsi = imsi
	client.addr = addr

	c.mutex.Lock()
	c.clientList.PushBack(client)
	c.imsi[imsi] = client
	c.mutex.Unlock()

	c.print()

	return true
}

func (c *clientMgr) init() {
	c.clientList = list.New()
	c.mutex = &sync.Mutex{}
	c.imsi = make(map[string]*clientInfo)
}

/*----------------------------------------------------------------------------*/

var cMgr clientMgr

func Init() {
	cMgr.init()
	go watcher()
}

func watcher() {
	for {
		time.Sleep(3 * time.Second)
		fmt.Println("[watcher] ... ing ...")
		cMgr.reArrange()
	}
}

func Insert(addr net.UDPAddr, imsi string, seq uint32) bool {
	return cMgr.insert(addr, imsi, seq)
}
