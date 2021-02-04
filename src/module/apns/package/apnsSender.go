package apnsHandler

import (
	"fmt"
	"log"
	"time"

	"github.com/anachronistic/apns"
)

type ApnsSender struct {
	ApnsGateway    string
	CertKeyPath    string
	ProcessorCount int

	deliverChan chan *apns.PushNotification
}

func (a *ApnsSender) processor(index int) {
	log.Printf("[apns/processor:%d] go\n", index)

	for {
		pn := <-a.deliverChan

		client := apns.NewClient(a.ApnsGateway, a.CertKeyPath, a.CertKeyPath)
		resp := client.Send(pn)
		if resp.Success == false {
			log.Printf("[apns/processor:%d] error: %v\n", index, resp.Error)
			continue
		}

		sendingMsg, _ := pn.PayloadString()

		log.Printf("[apns/processor:%d] Alert  : %v\n", index, sendingMsg)
		log.Printf("[apns/processor:%d] Success: %v\n", index, resp.Success)
		log.Printf("[apns/processor:%d] Error  : %v\n", index, resp.Error)
	}
}

func (a *ApnsSender) GenerateSend(deviceToken string, caller string, ip string, port string, alert string) error {
	if 0 >= len(deviceToken) || 0 >= len(caller) || 0 >= len(ip) || 0 >= len(port) {
		return fmt.Errorf("arguments check again")
	}

	payload := apns.NewPayload()
	payload.Alert = alert
	payload.Sound = "default"
	payload.Badge = 1

	pn := apns.NewPushNotification()
	pn.DeviceToken = deviceToken
	pn.AddPayload(payload)
	pn.Set("caller", caller)
	pn.Set("time", time.Now())
	pn.Set("ip", ip)
	pn.Set("port", port)

	a.deliverChan <- pn

	return nil
}

func (a *ApnsSender) Init(apnsGateway string, certKeyPath string, processorCount int) bool {
	if 0 >= len(apnsGateway) || 0 >= len(certKeyPath) {
		return false
	}
	a.ApnsGateway = apnsGateway
	a.CertKeyPath = certKeyPath
	a.ProcessorCount = processorCount
	a.deliverChan = make(chan *apns.PushNotification)

	for i := 0; i < a.ProcessorCount; i++ {
		go a.processor(i)
	}

	return true
}
