package goka

import "context"

type Key int

const (
	keySystem Key = 1 + iota
)

type Context interface {
	context.Context
	Name() string
	Path() string
	Sender() Actor
	Reply(data any) bool
	Send(target Actor, data any) bool
	SelectActor(path string) Actor
}

type ctx struct {
	context.Context
	cancel context.CancelFunc

	system ActorSystem
	actor  Actor
	sender Actor
}

func (c *ctx) Name() string {
	return c.actor.Name()
}

func (c *ctx) Path() string {
	return c.actor.Path()
}

func (c *ctx) Sender() Actor {
	return c.sender
}

func (c *ctx) Reply(data any) bool {
	return c.Send(c.sender, data)
}

func (c *ctx) Send(target Actor, data any) bool {
	return c.actor.Send(target, data)
}
func (c *ctx) SelectActor(path string) Actor {
	return c.system.SelectActor(path)
}
