package models

type Contact struct {
	user User
	from string
}

func (c Contact) From() string {
	return c.from
}

func NewContact(user User, from string) Contact {
	return Contact{user: user, from: from}
}
