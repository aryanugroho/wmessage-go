package pkg

// Message is message data model
type Message struct {
	Text string
}

type MessageDomain struct {
	store MessageStore
}

func (d MessageDomain) CreateMessage(message *Message) error {
	_, err := d.store.Add(message)
	if err != nil {
		return err
	}

	return nil
}

// MessageStore is message related pkg
type MessageStoreItf interface {
	Add(model *Message) (*Message, error)
	GetAll() ([]*Message, error)
}

type MessageStore struct {
}

func (m MessageStore) Add(model *Message) (*Message, error) {
	panic("implement me")
}

func (m MessageStore) GetAll() ([]*Message, error) {
	panic("implement me")
}
