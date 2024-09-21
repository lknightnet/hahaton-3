package model

import "fmt"

type Mail struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Topic   string   `json:"topic"`
	Message string   `json:"message"`
}

func (m *Mail) NewMessage(topic, message string) []byte {
	str := fmt.Sprintf("Subject: %s\r\n", topic)
	str += fmt.Sprintf("\n%s\r\n", message)
	return []byte(str)
}
