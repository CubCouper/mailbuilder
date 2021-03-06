Simple library for generating e-mails to send with golang's net/smtp.

[![GoDoc](https://godoc.org/github.com/zerobfd/mailbuilder?status.png)](https://godoc.org/github.com/zerobfd/mailbuilder)
[![Build Status](https://travis-ci.org/zerobfd/mailbuilder.png?branch=master)](https://travis-ci.org/zerobfd/mailbuilder)

Supports
------------
* Arbitrary nesting of multiparts
* Delayed building of e-mail body until message.Bytes() is called
* Implements json.UnmarshalJSON, see example.json and examplemulti.json for format

Todo
------------
* None planned, accepting issues

Not planned
------------
* Validation (multiparts can only use certain encoding, etc.)

Example use here:

```golang
package main

import(
        "github.com/zerobfd/mailbuilder"
        "fmt"
        "net/smtp"
      )

func main() {
  message := mailbuilder.NewMessage()
  message.AddTo(mailbuilder.NewAddress("recip@example.net", "Recipient"))
  message.From = mailbuilder.NewAddress("sender@example.net", "Sender")
  message.Subject = "Subject"
  body := mailbuilder.NewSimplePart()
  message.SetBody(body)
  body.AddHeader("Content-Type", "text/plain; charset=utf8")
  body.AddHeader("Content-Transfer-Encoding", "quoted-printable")
  body.Content = "Hello from golang!\n"
  auth := //auth info
  err := smtp.SendMail("smtp.example.com:587",
                auth,
                message.From.Email,
                message.Recipients(),
                message.Bytes())
  if (err != nil) {fmt.Printf("%v", err)}
}
```
