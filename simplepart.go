package mailbuilder

import (
	"bytes"
)

//Represents a simple mail body.
type SimplePart struct {
	Content string
	headers map[string]string
}

//Returns a pointer to a new Simplepart.
func NewSimplePart() *SimplePart {
	return &SimplePart{"", make(map[string]string, 0)}
}

//Adds a header to the part.
func (self *SimplePart) AddHeader(key, value string) {
	self.headers[key] = value
}

//Returns a function that builds the body when called.
//Shouldn't need to call directly, message.Bytes() does this for you.
func (self *SimplePart) Bytes() func() []byte {
	return func() []byte { return self.bytes() }
}

//Builds the body of the simplepart.
func (self *SimplePart) bytes() []byte {
	var b bytes.Buffer
	for k, v := range self.headers {
		b.WriteString(k + ": " + v + "\n")
	}
	b.WriteString("\n")
	b.WriteString(self.Content)

	return b.Bytes()
}
