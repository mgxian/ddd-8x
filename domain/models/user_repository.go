package models

type UserRepository interface {
	FindById(id int) User

	InSubscribeContext() SubscribeContext
	InOrderContext() OrderContext
	InSocialContext() SocialContext
}
