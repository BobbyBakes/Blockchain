package service

var MsgService msgServiceInterface

type msgService struct{}

type msgServiceInterface interface {
	CreatePhrase() bool
}

func (ms *msgService) CreatePhrase() bool { return true }

func init() {
	MsgService = &msgService{}
}
