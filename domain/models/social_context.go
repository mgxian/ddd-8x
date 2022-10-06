package models

type SocialContext interface {
	AsContact(user User) Contact
}
