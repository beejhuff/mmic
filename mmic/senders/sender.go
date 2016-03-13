package senders

type Sender interface {
	Send(payload interface{}) error
}
