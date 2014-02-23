package mailbuilder

import (
	"bytes"
	"strings"
)

type Message struct {
	body    func() []byte
	headers map[string]string
	To      []*Address
	Cc      []*Address
	Bcc     []*Address
	From    *Address
	Subject string
}

func NewMessage() *Message {
	message := new(Message)
	message.To = make([]*Address, 0)
	message.Cc = make([]*Address, 0)
	message.Bcc = make([]*Address, 0)
	return message
}

func (self *Message) AddTo(address *Address) {
	self.To = append(self.To, address)
}

func (self *Message) AddCc(address *Address) {
	self.Cc = append(self.Cc, address)
}

func (self *Message) AddBcc(address *Address) {
	self.Bcc = append(self.Bcc, address)
}

func (self *Message) SetBody(part BodyPart) {
	self.body = part.Bytes()
}

func (self *Message) AddHeader(key, value string) {
	self.headers[key] = value
}

func formatAddresses(adds []*Address) string {
	result := make([]string, 0, len(adds))
	for _, a := range adds {
		result = append(result, a.Full())
	}
	return strings.Join(result, ", ")
}

func (self *Message) Recipients() []string {
	result := make([]string, 0, len(self.To)+len(self.Cc)+len(self.Bcc))
	for _, a := range self.To {
		result = append(result, a.Email)
	}
	for _, a := range self.Cc {
		result = append(result, a.Email)
	}
	for _, a := range self.Bcc {
		result = append(result, a.Email)
	}
	return result
}

func (self *Message) Bytes() []byte {
	var b bytes.Buffer
	b.WriteString("From: ")
	b.WriteString(self.From.Full())
	b.WriteString("\nTo: ")
	b.WriteString(formatAddresses(self.To))
	if len(self.Cc) > 0 {
		b.WriteString("\nCc: ")
		b.WriteString(formatAddresses(self.Cc))
	}
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
