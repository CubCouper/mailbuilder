package mailbuilder

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
)

func TestUnmarshalJSON_simple(t *testing.T) {
	//set up the correct message
	message := NewMessage()
	message.From = NewAddress("test@test.com", "Testing")
	message.AddTo(NewAddress("testing@testing.com", "Test"))
	message.Subject = "Test Subject"
	text := NewSimplePart()
	text.Content = "Testing!"
	text.AddHeader("Content-Type", "text/plain")
	text.AddHeader("Content-Transfer-Encoding", "quoted-printable")
	message.SetBody(text)

	fromjson := &Message{}
	file, err := os.Open("example.json")
	if err != nil {
		t.Errorf("Couldn't open example json file!")
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&fromjson)
	if err != nil {
		t.Errorf("Error decoding JSON!")
		return
	}

	if message.Subject != fromjson.Subject {
		t.Errorf("Subject is not the same! Expected %v got %v", message.Subject, fromjson.Subject)
	}
	if message.From.Name != fromjson.From.Name {
		t.Errorf("From address is not the same! Expected %v got %v", message.From.Name, fromjson.From.Name)
	}
	if message.From.Email != fromjson.From.Email {
		t.Errorf("From address is not the same! Expected %v got %v", message.From.Email, fromjson.From.Email)
	}

	for i := range message.To {
		if message.To[i].Name != fromjson.To[i].Name {
			t.Errorf("To address is not the same! Expected %v got %v", message.To[i].Name, fromjson.To[i].Name)
		}
		if message.To[i].Email != fromjson.To[i].Email {
			t.Errorf("To address is not the same! Expected %v got %v", message.To[i].Email, fromjson.To[i].Email)
		}
	}

	if !bytes.Equal(message.Bytes(), fromjson.Bytes()) {
		t.Errorf("Body content is different!")
	}
}

func TestUnmarshalJSON_multi(t *testing.T) {
	//set up the correct message
	message := NewMessage()
	message.From = NewAddress("test@test.com", "Testing")
	message.AddTo(NewAddress("testing@testing.com", "Test"))
	message.Subject = "Test Subject"
	text := NewSimplePart()
	text.Content = "Testing!"
	text.AddHeader("Content-Type", "text/plain")
	text.AddHeader("Content-Transfer-Encoding", "quoted-printable")
	html := NewSimplePart()
	html.Content = "<h1>Testing!</h1>"
	html.AddHeader("Content-Type", "text/html")
	html.AddHeader("Content-Transfer-Encoding", "quoted-printable")
	alternative := NewMultiPart("multipart/alternative")
	alternative.AddPart(text)
	alternative.AddPart(html)
	message.SetBody(alternative)

	fromjson := &Message{}
	file, err := os.Open("examplemulti.json")
	if err != nil {
		t.Errorf("Couldn't open example json file!")
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&fromjson)
	if err != nil {
		t.Errorf("Error decoding JSON!")
		return
	}

	if message.Subject != fromjson.Subject {
		t.Errorf("Subject is not the same! Expected %v got %v", message.Subject, fromjson.Subject)
	}
	if message.From.Name != fromjson.From.Name {
		t.Errorf("From address is not the same! Expected %v got %v", message.From.Name, fromjson.From.Name)
	}
	if message.From.Email != fromjson.From.Email {
		t.Errorf("From address is not the same! Expected %v got %v", message.From.Email, fromjson.From.Email)
	}

	for i := range message.To {
		if message.To[i].Name != fromjson.To[i].Name {
			t.Errorf("To address is not the same! Expected %v got %v", message.To[i].Name, fromjson.To[i].Name)
		}
		if message.To[i].Email != fromjson.To[i].Email {
			t.Errorf("To address is not the same! Expected %v got %v", message.To[i].Email, fromjson.To[i].Email)
		}
	}

	//We can't compare Bytes directly because we expect the boundaries for multiparts to differ
	expectedlines := bytes.Split([]byte("\n"), message.Bytes())
	actuallines := bytes.Split([]byte("\n"), fromjson.Bytes())
	for i := range expectedlines {
		expectedline := expectedlines[i]
		actualline := actuallines[i]
		if !bytes.Contains([]byte("--==_Part_"), expectedline) {
			if !bytes.Equal(expectedline, actualline) {
				t.Errorf("Body content is different!")
			}
		}
	}
}
