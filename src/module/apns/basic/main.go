package main

import (
	"fmt"
	"time"

	"github.com/anachronistic/apns"
)

const (
	apnsGateway = "gateway.sandbox.push.apple.com:2195"
	deviceToken = "abababababadfadfadfadfadfd"
	certKeyPath = "../cert/voip_services_certificate.pem"
)

func main() {
	payload := apns.NewPayload()
	payload.Alert = "hello world"
	payload.Sound = "default"
	payload.Badge = 42

	pn := apns.NewPushNotification()
	pn.DeviceToken = deviceToken
	pn.AddPayload(payload)
	pn.Set("caller", "01012341234")
	pn.Set("time", time.Now())
	pn.Set("ip", "127.0.0.1")
	pn.Set("port", "0")
	alert, _ := pn.PayloadString()

	client := apns.NewClient(apnsGateway, certKeyPath, certKeyPath)
	resp := client.Send(pn)

	fmt.Println("Alert  : ", alert)
	fmt.Println("Success: ", resp.Success)
	fmt.Println("Error  : ", resp.Error)
}
