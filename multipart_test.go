package mailbuilder

import "testing"

func TestBoundaries(t *testing.T) {
	m1 := NewMultiPart("")
	m2 := NewMultiPart("")
	if m1.boundary == m2.boundary {
		t.Errorf("Boundaries are same: %v, %v", m1.boundary, m2.boundary)
	}
}

func ExampleMultiPart_s() {
	multi := NewMultiPart("multipart/alternative")
	html := NewSimplePart()
	text := NewSimplePart()

	//add content/headers to html and text

	multi.AddPart(text)
	multi.AddPart(html)

	message := NewMessage()
	message.SetBody(multi)
}

func ExampleMultiPart_c() {
	//nesting multiparts, with the mixed as top level
	alt := NewMultiPart("multipart/alternative")
	mix := NewMultiPart("multipart/mixed")
	mix.AddPart(alt)
	html := NewSimplePart()
	text := NewSimplePart()

	//add content/headers to html and text

	alt.AddPart(text)
	alt.AddPart(html)

	pic := NewSimplePart()
	pic.AddHeader("Content-Type", "image/jpeg")
	pic.AddHeader("Content-Transfer-Encoding", "base64")
	pic.Content = "==BASE64DATA"

	mix.AddPart(pic)

	message := NewMessage()
	message.SetBody(mix)
}
