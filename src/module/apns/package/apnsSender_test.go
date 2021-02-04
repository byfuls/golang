package apnsHandler

import (
	"log"
	"testing"
)

const (
	apnsGateway = "gateway.sandbox.push.apple.com:2195"
	certKeyPath = "../cert/voip_services_certificate.pem"

	deviceToken = "eaefa10820148fdc1495b86491f7cfeda518cde7c16914d123e701920daa5a82"
	caller      = "01012341234"
	ip          = "127.0.0.1"
	port        = "3000"
	alert       = "alertMsg"
)

func TestApnsSend(t *testing.T) {
	apnsHandler := ApnsSender{}

	if !apnsHandler.Init(apnsGateway, certKeyPath, 3) {
		log.Println("init error")
		return
	}
	if err := apnsHandler.GenerateSend(deviceToken, caller, ip, port, alert); err != nil {
		log.Println("generateSend error: ", err)
		return
	} else {
		log.Println("send ok")
	}
}
