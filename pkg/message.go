package pkg

// Message is message data model
type Message struct {
	Text string `json:"text"`
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

func (d MessageDomain) GetAll() ([]*Message, error) {
	list, err := d.store.GetAll()
	if err != nil {
		return nil, err
	}

	return list, nil
}

var messageStore *MessageStore

// getMessageStore is singleton accessor for messageStore
func getMessageStore() MessageStoreItf {
	if messageStore == nil {
		messageStore = &MessageStore{}
	}
	return messageStore
}

// MessageStoreItf is message related pkg contract
type MessageStoreItf interface {
	Add(model *Message) (*Message, error)
	GetAll() ([]*Message, error)
}

// messageData will act like data store
var messageData []*Message

// MessageStore is message related pkg
type MessageStore struct {
}

func (m MessageStore) Add(model *Message) (*Message, error) {
	messageData = append(messageData, model)

	return model, nil
}

func (m MessageStore) GetAll() ([]*Message, error) {
	return messageData, nil
}
