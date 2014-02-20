package mailbuilder

import (
	"bytes"
	"strconv"
  "math/rand"
)

type MultiPart struct {
	content               []func() []byte
	headers               map[string]string
	boundary, contentType string
}

func NewMultiPart(ctype string) *MultiPart {
	bound := "--==_Part_" + strconv.Itoa(rand.Int()) + "=="
	return &MultiPart{make([]func() []byte, 0),
		make(map[string]string, 0),
		bound,
		ctype}
}

func (self *MultiPart) AddPart(part BodyPart) {
	self.content = append(self.content, part.Bytes())
}

func (self *MultiPart) SetContentType(ctype string) {
	self.contentType = ctype
}

func (self *MultiPart) AddHeader(key, value string) {
	self.headers[key] = value
}

func (self *MultiPart) Bytes() func() []byte {
	return func() []byte { return self.bytes() }
}

func (self *MultiPart) bytes() []byte {
	var b bytes.Buffer
	b.WriteString("Content-Type: " + self.contentType + ";\n\tboundary=\"" + self.boundary + "\"")
	for k, v := range self.headers {
		b.WriteString(k + ": " + v + "\n")
	}
	b.WriteString("\n\nThis is a multi-part message in MIME format.\n\n")
	for _, f := range self.content {
		b.WriteString("--" + self.boundary + "\n")
		b.Write(f())
		b.WriteString("\n")
	}
	b.WriteString("--" + self.boundary + "--\n")

	return b.Bytes()
}
