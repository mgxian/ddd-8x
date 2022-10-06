package models

type OrderContext interface {
	AsBuyer(user User) Buyer
}
