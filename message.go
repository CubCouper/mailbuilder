//Simple library for building messages to send with net/smtp.
//See github.com/zerobfd/mailbuilder
package mailbuilder

import (
	"bytes"
	"strings"
)

//The main struct in mailbuilder.
//Contains info on from, recipients, subject, etc.
//Body is a func() []byte to delay actually building until Bytes() is called.
//This allows attaching/adding info in whatever order you like.
type Message struct {
	body    BodyPart
	headers map[string]string
	To      []*Address
	Cc      []*Address
	Bcc     []*Address
	From    *Address
	Subject string
}

//Returns pointer to a new message
//Initializes to, cc, and bcc so AddTo etc. don't error
func NewMessage() *Message {
	message := new(Message)
	message.To = make([]*Address, 0)
	message.Cc = make([]*Address, 0)
	message.Bcc = make([]*Address, 0)
	return message
}

//Add a new Address to the to list
func (self *Message) AddTo(address *Address) {
	self.To = append(self.To, address)
}

//Add a new Address to the cc list
func (self *Message) AddCc(address *Address) {
	self.Cc = append(self.Cc, address)
}

//Add a new Address to the bcc list
func (self *Message) AddBcc(address *Address) {
	self.Bcc = append(self.Bcc, address)
}

//Set the main body part of the message.
//Takes either a simplepart or multipart.
func (self *Message) SetBody(part BodyPart) {
	self.body = part
}

//Sets a MIME header to specified value
func (self *Message) AddHeader(key, value string) {
	self.headers[key] = value
}

//Used in constructing the MIME to/cc headers to ensure format is correct
func formatAddresses(adds []*Address) string {
	result := make([]string, 0, len(adds))
	for _, a := range adds {
		result = append(result, a.Full())
	}
	return strings.Join(result, ", ")
}

//Returns all recipients' email addresses, for SMTP use
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

//Constructs the mail body after everything has been added
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
	b.Write(self.body.Bytes())

	return b.Bytes()
}
