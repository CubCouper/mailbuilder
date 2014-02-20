package mailbuilder

type Address struct {
	Name, Email string
}

func NewAddress(email, name string) *Address {
	return &Address{name, email}
}

func (self *Address) Full() string {
	return self.Name + " <" + self.Email + ">"
}
