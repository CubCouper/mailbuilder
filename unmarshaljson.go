package mailbuilder

import (
	"encoding/json"
)

//This struct is a temporary holder while we figure out if we
//need a multipart or not.
type JSONEmail struct {
	Body    JSONBody          `json:"body"`
	Headers map[string]string `json:"headers"`
	To      []*Address        `json:"to"`
	Cc      []*Address        `json:"cc"`
	Bcc     []*Address        `json:"bcc"`
	From    *Address          `json:"from"`
	Subject string            `json:"subject"`
}

//This struct is a temporary holder for a BodyPart while we figure
//out whether it's a MultiPart or SimplePart
type JSONBody struct {
	Multipart string          `json:"multipart"`
	Message   json.RawMessage `json:"message"`
}

//Implement json.Unmarshaller for Message
func (mess *Message) UnmarshalJSON(data []byte) error {
	jsonemail := &JSONEmail{}
	err := json.Unmarshal(data, &jsonemail)
	if err != nil {
		return err
	}
	if jsonemail.Body.Multipart == "no" {
		body := NewSimplePart()
		err = json.Unmarshal(jsonemail.Body.Message, &body)
		if err != nil {
			return err
		}
		mess.SetBody(body)
	} else {
		body := NewMultiPart("multipart/" + jsonemail.Body.Multipart)
		err = json.Unmarshal(jsonemail.Body.Message, &body)
		if err != nil {
			return err
		}
		mess.SetBody(body)
	}
	mess.From = jsonemail.From
	mess.headers = jsonemail.Headers
	mess.To = jsonemail.To
	mess.Cc = jsonemail.Cc
	mess.Bcc = jsonemail.Bcc
	mess.Subject = jsonemail.Subject
	return nil
}

//Implement json.Unmarshaller for MultiPart
func (multi *MultiPart) UnmarshalJSON(data []byte) error {
	var jsonbody []JSONBody
	err := json.Unmarshal(data, &jsonbody)
	if err != nil {
		return err
	}
	for _, v := range jsonbody {
		if v.Multipart == "no" {
			body := NewSimplePart()
			err = json.Unmarshal(v.Message, &body)
			if err != nil {
				return err
			}
			multi.AddPart(body)
		} else {
			body := NewMultiPart("multipart/" + v.Multipart)
			err = json.Unmarshal(v.Message, &body)
			if err != nil {
				return err
			}
			multi.AddPart(body)
		}
	}
	return nil
}
