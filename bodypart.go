package mailbuilder

type BodyPart interface {
	AddHeader(key, value string)
	Bytes() []byte
}
