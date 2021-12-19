package goka

import "context"

type Message struct {
	_      struct{}
	data   any
	sender Actor
	next   *Message
}

type Messages struct {
	head  *Message
	tail  *Message
	count int
}

func (m *Messages) Length() int {
	return m.count
}

func (m *Messages) Add(msg *Message) {
	if m.head == nil {
		m.head = msg
		m.tail = msg
	} else {
		m.tail.next = msg
		m.tail = msg
	}
	m.count++
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
	m.count--
	return msg
}

type ActorRef interface {
	Receive(sender Actor, data any)
}

type Actor interface {
	Name() string
	Path() string
	Tell(ctx context.Context, target Actor, data any) bool
	accept(ctx context.Context, msg *Message) bool
}

type ActorSystem interface {
	Name() string
	ActorOf(name string, ref ActorRef) Actor
}
