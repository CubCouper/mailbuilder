Simple library for generating e-mails to send with golang's net/smtp.

Supports
------------
* Arbitrary nesting of multiparts
* Delayed building of e-mail body until message.Bytes() is called

Todo
------------
* Add some functionality for message recipients
* Tests
* Comments
* Godoc

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
  message := new(mailbuilder.Message)
  message.To = mailbuilder.NewAddress("recip@example.net", "Recipient")
  message.From = mailbuilder.NewAddress("sender@example.net", "Sender")
  message.Subject = "Subject"
  body := mailbuilder.NewSimplePart()
  message.SetBody(body)
  body.AddHeader("Content-Type", "text/plain; charset=utf8")
  body.AddHeader("Content-Transfer-Encoding", "quoted-printable")
  body.Content = "Hello from golang!\n"
  auth := //gmail information
  err := smtp.SendMail("smtp.gmail.com:587",
                auth,
                message.From.Email,
                []string{message.To.Email},
                message.Bytes())
  if (err != nil) {fmt.Printf("%v", err)}
}
```
