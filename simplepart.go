package mailbuilder

import (
	"bytes"
)

//Represents a simple mail body.
type SimplePart struct {
	Content string            `json:"content"`
	Headers map[string]string `json:"headers"`
}

//Returns a pointer to a new Simplepart.
func NewSimplePart() *SimplePart {
	return &SimplePart{"", make(map[string]string, 0)}
}

//Adds a header to the part.
func (self *SimplePart) AddHeader(key, value string) {
	self.Headers[key] = value
}

//Builds the body of the simplepart.
func (self *SimplePart) Bytes() []byte {
	var b bytes.Buffer
	for k, v := range self.Headers {
		b.WriteString(k + ": " + v + "\n")
	}
	b.WriteString("\n")
	b.WriteString(self.Content)

	return b.Bytes()
}
