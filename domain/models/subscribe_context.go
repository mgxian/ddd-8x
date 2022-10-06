package models

type SubscribeContext interface {
	AsReader(user User) Reader
}
