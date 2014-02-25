package mailbuilder

func ExampleSimplePart() {
	text := NewSimplePart()
	text.AddHeader("Content-Type", "text/plain; charset=utf8")
	text.AddHeader("Content-Transfer-Encoding", "quoted-printable")
	text.Content = "Hello from golang!"

	message := NewMessage()
	message.SetBody(text)
}
