package service

import "github.com/aryanugroho/wmessage-go/pkg"

func AddMessage(message *pkg.Message) error {
	messageDomain := pkg.GetMessageDomain()
	err := messageDomain.CreateMessage(message)
	if err != nil {
		return err
	}

	return nil
}

func GetMessages() ([]*pkg.Message, error) {
	messageDomain := pkg.GetMessageDomain()
	list, err := messageDomain.GetAll()
	if err != nil {
		return nil, err
	}
	return list, err
}
