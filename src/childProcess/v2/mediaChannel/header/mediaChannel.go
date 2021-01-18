package mediaChannel

import (
	"encoding/json"
)

type MediaMessage struct {
	Head string `json:"head"`
	Body string `json:"body"`
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

func GenerateMediaMessage(head string, body string, ip string, port string) MediaMessage {
	return MediaMessage{
		Head: head,
		Body: body,
		Ip:   ip,
		Port: port,
	}
}

func ParsingMediaMessage(message []byte) MediaMessage {
	var mediaMessage MediaMessage
	err := json.Unmarshal(message, &mediaMessage)
	if err != nil {
		return MediaMessage{}
	}
	return mediaMessage
}
