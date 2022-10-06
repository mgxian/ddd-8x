package models

type Reader struct {
	user            User
	from            string
	mySubscriptions MySubscriptions
}

func (r Reader) MySubscriptions() MySubscriptions {
	return r.mySubscriptions
}

func (r Reader) SetMySubscriptions(mySubscriptions MySubscriptions) Reader {
	r.mySubscriptions = mySubscriptions
	return r
}

func (r Reader) From() string {
	return r.from
}

func NewReader(user User, from string) Reader {
	return Reader{user: user, from: from}
}
