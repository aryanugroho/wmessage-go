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
	// messageData will act like data store
	messageData []*Message
}

func (m MessageStore) Add(model *Message) (*Message, error) {
	m.messageData = append(m.messageData, model)

	return model, nil
}

func (m MessageStore) GetAll() ([]*Message, error) {
	panic("implement me")
}
