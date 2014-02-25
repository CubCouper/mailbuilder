package mailbuilder

import "net/smtp"

func ExampleMessage() {
	text := NewSimplePart()

	message := NewMessage()
	message.SetBody(text)

	message.AddTo(NewAddress("recip1@example.net", "Recipient"))
	message.AddCc(NewAddress("recip2@example.net", "Recipient's Brother"))
	message.AddBcc(NewAddress("recip3@example.net", "Secret Recipient"))
	message.From = NewAddress("someguy@example.ch", "Recipient's Acquaintance")

	err := smtp.SendMail("address:port",
		nil, //some auth mechanism
		message.From.Email,
		message.Recipients(),
		message.Bytes())

	if err != nil {
		//log
	}
}
