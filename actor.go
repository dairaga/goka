package goka

type ActorRef interface {
	Receive(ctx Ctx, data any)
}

type Actor interface {
	Name() string
	Path() string
	Send(target Actor, data any) bool
	accept(msg *Message) bool
}

type actor struct {
	_ struct{}

	ctx  Ctx
	name string
	path string
	ref  ActorRef

	closed bool
	inbox  chan *Message
	done   chan struct{}
	job    chan *Message

	messages *Messages
}

var _ Actor = &actor{}

func (a *actor) Name() string {
	return a.name
}

func (a *actor) Path() string {
	return a.path
}

func (a *actor) Send(target Actor, data any) bool {
	if target == nil {
		return false
	}

	msg := &Message{
		data:   data,
		sender: a,
		next:   nil,
	}
	return target.accept(msg)
}

func (a *actor) accept(msg *Message) bool {
	if a.closed {
		return false
	}

	select {
	case a.inbox <- msg:
		return true
	case <-a.ctx.Done():
		return false
	}
}

func (a *actor) run() {
	for {
		select {
		case msg := <-a.inbox:
			a.messages.Add(msg)
		case _, ok := <-a.done:
			if !ok {
				return
			}
			msg := a.messages.Get()
			if msg != nil {
				a.job <- msg
			}
		}
	}
}

func (a *actor) exec() {
	for {
		msg, ok := <-a.job
		if !ok || msg == nil {
			return
		}
		switch v := msg.data.(type) {
		case Identify:

		default:
			a.ref.Receive(a.ctx.setSender(msg.sender), v)
		}

		a.done <- struct{}{}
	}
}
