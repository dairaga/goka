package goka

import "context"

type Key int

const (
	keySystem Key = 1 + iota
)

type Ctx interface {
	context.Context
	Reply(data any) bool
	Send(taret Actor, data any) bool
	Sender() Actor

	ActorOf(name string, ref ActorRef) Actor
	SelectActor(path string) Actor

	setSender(sender Actor) Ctx
}

type ctx struct {
	context.Context
	sender Actor
	actor  Actor
	system ActorSystem
}

var _ Ctx = &ctx{}

func (c *ctx) setSender(sender Actor) Ctx {
	c.sender = sender
	return c
}

func (c *ctx) Sender() Actor {
	return c.sender
}

func (c *ctx) Reply(data any) bool {
	return false
}

func (c *ctx) Send(taret Actor, data any) bool {
	return false
}

func (c *ctx) ActorOf(name string, ref ActorRef) Actor {
	return nil
}

func (c *ctx) SelectActor(path string) Actor {
	return nil
}
