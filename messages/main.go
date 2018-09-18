package messages

// Messages represent a collection of Message
type Messages []Message

// Message it is a simple structure that keep error message info
type Message struct {
	Type  string
	Name  string
	Value string
}

// Get will return all collection of Messages
func (m Messages) Get() Messages {
	return m
}

// GetCount return number of Message into Messages collection
func (m Messages) GetCount() int {
	return len(m)
}

// AddError will create Message structure and it into Messages collection
func (m *Messages) AddError(v, n string) {
	msg := Message{}
	msg.Type = "error"
	msg.Name = n
	msg.Value = v
	*m = append(*m, msg)
}
