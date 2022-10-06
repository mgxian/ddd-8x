package models

type Buyer struct {
	user User
	from string
}

func NewBuyer(user User, from string) Buyer {
	return Buyer{user: user, from: from}
}

func (b Buyer) From() string {
	return b.from
}
