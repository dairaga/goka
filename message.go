package goka

type Message struct {
	_      struct{}
	sender Actor
	data   any
	next   *Message
}

func (m *Message) Raw() any {
	return m.data
}

func (m *Message) Sender() Actor {
	return m.sender
}

type Messages struct {
	_      struct{}
	length int
	head   *Message
	tail   *Message
}

func (m *Messages) Add(msg *Message) {
	if m.head == nil {
		m.head = msg
		m.tail = msg
	} else {
		m.tail.next = msg
		m.tail = msg
	}
	m.length++
}

func (m *Messages) Get() *Message {
	if m.head == nil {
		return nil
	}
	msg := m.head
	m.head = msg.next
	if m.head == nil {
		m.tail = nil
	}
	m.length--
	return msg
}

func (m *Messages) Length() int {
	return m.length
}
