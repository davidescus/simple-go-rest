package messages

import "fmt"

// Messages represent a collection of Message
type Messages []Message

type Messenger interface {
	Get() Messages
	GetCount() int
	AddError(string, string)
}

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
func (m *Messages) AddError(name, value string) {

	fmt.Println("here")

	if value != "" {

		fmt.Println(name + " " + value)

		msg := Message{}
		msg.Type = "error"
		msg.Name = name
		msg.Value = value
		*m = append(*m, msg)
	}
}
