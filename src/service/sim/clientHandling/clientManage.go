package clientHandling

import (
	"container/list"
	"net"
	"sync"
	"time"

	"module/logging"
)

type clientInfo struct {
	updatedTime time.Time
	imsi        string
	addr        net.UDPAddr

	gateway     string
	gatewayAddr net.UDPAddr
}

type clientMgr struct {
	clientList *list.List
	mutex      *sync.Mutex

	imsi    map[string]*clientInfo
	gateway map[string]*clientInfo
}

const (
	TIMEOUT = 120
)

func (c *clientMgr) reArrange() {
	now := time.Now()
	defer c.mutex.Unlock()
	c.mutex.Lock()
	
	idx := 0
	tcnt := c.clientList.Len()
	for e := c.clientList.Front(); e != nil; e = e.Next() {
		v := e.Value.(*clientInfo)
		elapsed := now.Sub(v.updatedTime)

		logging.DebugLn("[clientMgr/reArrange] [", idx+1, "/", tcnt, "] time: ", v.updatedTime)
		logging.DebugLn("[clientMgr/reArrange] [", idx+1, "/", tcnt, "] elapsed: ", elapsed)
		logging.DebugLn("[clientMgr/reArrange] [", idx+1, "/", tcnt, "] gatewayAddr: ", v.gatewayAddr)
		logging.DebugLn("[clientMgr/reArrange] [", idx+1, "/", tcnt, "] addr: ", v.addr)
		logging.DebugLn("[clientMgr/reArrange] [", idx+1, "/", tcnt, "] imsi: ", v.imsi)

		if int(elapsed.Seconds()) > TIMEOUT {
			logging.TraceF("[clientMgr/reArrange] [%d/%d] TIMEOUT", idx+1, tcnt)
			c.remove(e, v.imsi)
		}
		idx++
	}
}

func (c *clientMgr) print() {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	for e := c.clientList.Front(); e != nil; e = e.Next() {
		v := e.Value.(*clientInfo)

		logging.TraceLn("[clientMgr/reArrange] time: ", v.updatedTime)
		logging.TraceLn("[clientMgr/reArrange] gatewayAddr: ", v.gatewayAddr)
		logging.TraceLn("[clientMgr/reArrange] addr: ", v.addr)
		logging.TraceLn("[clientMgr/reArrange] imsi: ", v.imsi)
	}

	for key, element := range c.imsi {
		logging.TraceLn("Key: ", key, " => ", "Element: ", element)
		logging.TraceLn("updated time: ", element.updatedTime)
		logging.TraceLn("updated imsi: ", element.imsi)
	}
}

func (c *clientMgr) update(imsi string) bool {
	if 0 >= len(imsi) {
		logging.ErrorLn("[update-err] imsi: ", imsi)
		return false
	}

	if client := c.imsi[imsi]; client != nil {
		client.updatedTime = time.Now()
		logging.DebugLn("[update] complete, imsi: ", imsi)
		return true
	} else {
		logging.ErrorLn("[update-err] not found imsi: ", imsi)
		return false
	}
}

func (c *clientMgr) remove(element *list.Element, imsi string) {
	delete(c.imsi, imsi)
	c.clientList.Remove(element)
}

func (c *clientMgr) findViaGateway(gateway string) *clientInfo {
	if 0 >= len(gateway) {
		logging.ErrorLn("[findViaGateway-err] gateway: ", gateway)
		return nil
	}

	client := c.gateway[gateway]
	if client == nil {
		logging.ErrorLn("[findViaGateway-err] find error, gateway: ", gateway)
		return nil
	}

	return c.gateway[gateway]
}

func (c *clientMgr) findViaImsi(imsi string) *clientInfo {
	if 0 >= len(imsi) {
		logging.ErrorLn("[findViaImsi-err] imsi: ", imsi)
		return nil
	}

	return c.imsi[imsi]
}

func (c *clientMgr) mapping(addr net.UDPAddr, gateway string, imsi string) bool {
	if 0 >= len(gateway) || 0 >= len(imsi) {
		logging.ErrorLn("[mapping-err] gateway: ", gateway)
		return false
	}

	defer c.mutex.Unlock()
	c.mutex.Lock()
	client := c.findViaImsi(imsi)
	if client == nil {
		logging.ErrorLn("[mapping-err] not found client info via imsi: ", imsi)
		return false
	}
	client.gateway = gateway
	client.gatewayAddr = addr
	c.gateway[gateway] = client

	return true
}

func (c *clientMgr) insert(addr net.UDPAddr, imsi string, seq uint32) bool {
	if 0 >= len(imsi) {
		logging.ErrorLn("[insert-err] imsi: ", imsi)
		return false
	}

	client := new(clientInfo)
	client.updatedTime = time.Now()
	client.imsi = imsi
	client.addr = addr

	c.mutex.Lock()
	c.clientList.PushBack(client)
	c.imsi[imsi] = client
	c.mutex.Unlock()

	//c.print()

	return true
}

func (c *clientMgr) init() {
	c.clientList = list.New()
	c.mutex = &sync.Mutex{}
	c.imsi = make(map[string]*clientInfo)
	c.gateway = make(map[string]*clientInfo)
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
		logging.TraceLn("[watcher] re-arrange start")
		cMgr.reArrange()
	}
}

func Update(imsi string) bool {
	return cMgr.update(imsi)
}

func GetMappedAddressNGatewayViaImsi(imsi string) (bool, net.UDPAddr, string) {
	client := cMgr.findViaImsi(imsi)
	if client != nil {
		return true, client.addr, client.gateway
	}
	return false, client.addr, client.gateway
}

func GetMappedAddressNImsiViaGateway(gateway string) (bool, net.UDPAddr, string) {
	client := cMgr.findViaGateway(gateway)
	if client != nil {
		return true, client.addr, client.imsi
	}
	return false, client.addr, client.imsi
}

func Insert(addr net.UDPAddr, imsi string, seq uint32) bool {
	return cMgr.insert(addr, imsi, seq)
}

func Mapping(addr net.UDPAddr, gateway string, imsi string) bool {
	return cMgr.mapping(addr, gateway, imsi)
}
