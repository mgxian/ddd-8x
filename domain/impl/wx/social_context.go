package wx

import (
	"ddd/domain/models"
)

type SocialContext struct {
}

func (s SocialContext) AsContact(user models.User) models.Contact {
	return models.NewContact(user, "wx")
}

func NewSocialContext() SocialContext {
	return SocialContext{}
}
