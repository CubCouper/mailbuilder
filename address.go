package mailbuilder

//Simple struct that stores a name and an e-mail.
type Address struct {
	Name, Email string
}

//Returns a pointer to an Address with the specified name and e-mail
func NewAddress(email, name string) *Address {
	return &Address{name, email}
}

//Returns the name and e-mail for use in the MIME headers
func (self *Address) Full() string {
	return self.Name + " <" + self.Email + ">"
}
