package api

import (
	"ddd/domain/models"
)

type OrderContext struct {
}

func (o OrderContext) AsBuyer(user models.User) models.Buyer {
	return models.NewBuyer(user, "api")
}

func NewOrderContext() OrderContext {
	return OrderContext{}
}
