package pkg

// Message is message data model
type Message struct {
	Text string
}

// MessageDomain is message logic
type MessageDomain struct {
	store MessageStoreItf
}

func GetMessageDomain() *MessageDomain {
	return &MessageDomain{
		store: getMessageStore(),
	}
}

func (d MessageDomain) CreateMessage(message *Message) error {
	_, err := d.store.Add(message)
	if err != nil {
		return err
	}

	return nil
}

var messageStore *MessageStore

// getMessageStore is singleton accessor for messageStore
func getMessageStore() MessageStoreItf {
	if messageStore == nil {
		messageStore = &MessageStore{
			messageData: []*Message{},
		}
	}
	return messageStore
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
	return m.messageData, nil
}
