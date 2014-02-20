package mailbuilder

import (
	"bytes"
)

type Message struct {
	body    func() []byte
	headers map[string]string
	To      *Address
	From    *Address
	Subject string
}

func (self *Message) SetBody(part BodyPart) {
	self.body = part.Bytes()
}

func (self *Message) AddHeader(key, value string) {
	self.headers[key] = value
}

func (self *Message) Bytes() []byte {
	var b bytes.Buffer
	b.WriteString("From: ")
	b.WriteString(self.From.Full())
	b.WriteString("\nTo: ")
	b.WriteString(self.To.Full())
	b.WriteString("\nSubject: ")
	b.WriteString(self.Subject)
	b.WriteString("\n")
	for k, v := range self.headers {
		b.WriteString(k)
		b.WriteString(": ")
		b.WriteString(v)
		b.WriteString("\n")
	}
	b.Write(self.body())

	return b.Bytes()
}
