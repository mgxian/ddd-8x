package db

import "ddd/domain/models"

type MySubscriptions struct {
	user models.User
}

func NewMySubscriptions(user models.User) *MySubscriptions {
	return &MySubscriptions{user: user}
}

func (my MySubscriptions) SubList(from, to int) []models.Subscription {
	return nil
}

func (my MySubscriptions) Count() int {
	return 111
}
