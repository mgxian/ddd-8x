package models

type MySubscriptions interface {
	SubList(from, to int) []Subscription
	Count() int
}
