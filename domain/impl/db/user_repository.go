package db

import "ddd/domain/models"

type UserRepository struct {
	subscribeContext models.SubscribeContext
	socialContext    models.SocialContext
	orderContext     models.OrderContext
}

func (u UserRepository) InSubscribeContext() models.SubscribeContext {
	return u.subscribeContext
}

func (u UserRepository) InOrderContext() models.OrderContext {
	return u.orderContext
}

func (u UserRepository) InSocialContext() models.SocialContext {
	return u.socialContext
}

func NewUserRepository(subscribeContext models.SubscribeContext,
	socialContext models.SocialContext,
	orderContext models.OrderContext) *UserRepository {
	return &UserRepository{
		subscribeContext: subscribeContext,
		socialContext:    socialContext,
		orderContext:     orderContext,
	}
}

func (u UserRepository) FindById(id int) models.User {
	return models.NewUser("from-db")
}
