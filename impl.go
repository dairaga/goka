package goka

/*
type actor struct {
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

func (a *actor) Tell(ctx context.Context, target Actor, data any) bool {
	if target == nil {
		return false
	}
	return target.accept(ctx, &Message{data: data, sender: a})
}

func (a *actor) accept(ctx context.Context, msg *Message) bool {
	if a.closed {
		return false
	}

	select {
	case a.inbox <- msg:
		return true
	case <-ctx.Done():
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
		a.ref.Receive(msg.sender, msg.data)
		a.done <- struct{}{}
	}
}

func (a *actor) finalize() {
	if !a.closed {
		close(a.done)
		close(a.inbox)
	}
}

type system struct {
	name string
}

var _ ActorSystem = &system{}

func (s *system) Name() string {
	return s.name
}

func (s *system) ActorOf(name string, ref ActorRef) Actor {
	if ref == nil {
		panic("ref is nil")
	}

	ret := &actor{
		name:   name,
		path:   fmt.Sprintf("goka://%s/user/%s", s.name, name),
		ref:    ref,
		closed: false,
		inbox:  make(chan *Message, 1),
		job:    make(chan *Message),
		done:   make(chan struct{}),
		messages: &Messages{
			head:  nil,
			tail:  nil,
			count: 0,
		},
	}

	go ret.run()
	go ret.exec()

	return ret
}
*/
