package mailbuilder

import (
	"bytes"
)

type SimplePart struct {
	Content string
	headers map[string]string
}

func NewSimplePart() *SimplePart {
	return &SimplePart{"", make(map[string]string, 0)}
}

func (self *SimplePart) AddHeader(key, value string) {
	self.headers[key] = value
}

func (self *SimplePart) Bytes() func() []byte {
  return func() []byte {return self.bytes()}
}

func (self *SimplePart) bytes() []byte {
	var b bytes.Buffer
	for k, v := range self.headers {
		b.WriteString(k + ": " + v + "\n")
	}
	b.WriteString("\n")
	b.WriteString(self.Content)

	return b.Bytes()
}
