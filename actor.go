package goka

import (
	"context"
	"fmt"
)

type ActorRef interface {
	Receive(ctx Context, data any)
}

type Actor interface {
	Name() string
	Path() string

	Send(target Actor, data any) bool
	accept(msg *Message) bool
}

type actor struct {
	_ struct{}

	ctx *ctx

	ref ActorRef

	name string
	path string

	closed bool
	inbox  chan *Message
	done   chan struct{}
	job    chan *Message

	messages *Messages
}

func (a *actor) Name() string {
	return a.name
}

func (a *actor) Path() string {
	return a.path
}

func (a *actor) Send(target Actor, data any) bool {
	return target.accept(&Message{sender: a, data: data, next: nil})
}

func (a *actor) schedule() {
	msg := a.messages.Get()
	if msg != nil {
		a.job <- msg
	}
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
		case msg, ok := <-a.inbox:
			if ok && msg != nil {
				a.messages.Add(msg)
				if len(a.job) <= 0 {
					a.schedule()
				}
			}
		case _, ok := <-a.done:
			if ok {
				a.schedule()
			} else {
				// TODO: save unprocessed messages
				return
			}
		}
	}

}

func (a *actor) exec() {

	defer func() {
		a.closed = true
		close(a.job)
		close(a.done)
		close(a.inbox)
		a.ctx.cancel()
	}()

	go a.run()
	for {
		select {
		case msg, ok := <-a.job:

			if ok && msg != nil {
				a.ctx.sender = msg.sender
				a.ref.Receive(a.ctx, msg.data)
				a.done <- struct{}{}
			}
		case <-a.ctx.Done():
			return
		}
	}
}

func newActor(s *system, path, name string, ref ActorRef) *actor {
	if ref == nil {
		panic("ref is nil")
	}

	ctx := &ctx{
		system: s,
	}

	ctx.Context, ctx.cancel = context.WithCancel(s.ctx)

	ret := &actor{
		ctx:      ctx,
		ref:      ref,
		name:     name,
		path:     fmt.Sprintf("%s/%s", path, name),
		closed:   false,
		inbox:    make(chan *Message, 1),
		done:     make(chan struct{}, 1),
		job:      make(chan *Message, 1),
		messages: &Messages{},
	}

	ctx.actor = ret

	go ret.exec()

	return ret
}
