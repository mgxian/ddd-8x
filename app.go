package main

import (
	"ddd/domain/impl/api"
	"ddd/domain/impl/db"
	"ddd/domain/impl/wx"
	"ddd/domain/models"
	"fmt"
)

func main() {
	var users models.UserRepository

	users = db.NewUserRepository(db.NewSubscribeContext(), wx.NewSocialContext(), api.NewOrderContext())

	user := users.FindById(1)

	contact := users.InSocialContext().AsContact(user)
	fmt.Printf("contact from: [%s]\n", contact.From())

	buyer := users.InOrderContext().AsBuyer(user)
	fmt.Printf("buyer from: [%s]\n", buyer.From())

	reader := users.InSubscribeContext().AsReader(user)
	fmt.Printf("reader from: [%s]\n", reader.From())
	fmt.Printf("reader subscription count: [%d]\n", reader.MySubscriptions().Count())
}
