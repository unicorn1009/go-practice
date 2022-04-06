package myutil

type Queue interface {
	Offer(e interface{}) bool
	Peek() interface{}
	Poll() interface{}
}
