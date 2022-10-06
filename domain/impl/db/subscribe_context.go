package db

import (
	"ddd/domain/models"
)

type SubscribeContext struct {
}

func (s SubscribeContext) AsReader(user models.User) models.Reader {
	reader := models.NewReader(user, "db")
	return reader.SetMySubscriptions(NewMySubscriptions(user))
}

func NewSubscribeContext() SubscribeContext {
	return SubscribeContext{}
}
