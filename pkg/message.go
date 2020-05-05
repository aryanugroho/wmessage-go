package pkg

// Message is message data model
type Message struct {
	Text string
}

// MessageDomain is message logic
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

// MessageStoreItf is message related pkg contract
type MessageStoreItf interface {
	Add(model *Message) (*Message, error)
	GetAll() ([]*Message, error)
}

// MessageStore is message related pkg
type MessageStore struct {
}

func (m MessageStore) Add(model *Message) (*Message, error) {
	panic("implement me")
}

func (m MessageStore) GetAll() ([]*Message, error) {
	panic("implement me")
}
