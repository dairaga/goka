package goka

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

type ActorSystem interface {
	Name() string
	Path() string

	ActorOf(name string, ref ActorRef) Actor
	SelectActor(path string) Actor
	Wait()
	Shutdown()
}

type system struct {
	_ struct{}

	ctx    context.Context
	cancel context.CancelFunc

	name string
	path string

	actors map[string]Actor
}

func (s *system) Name() string {
	return s.name
}

func (s *system) Path() string {
	return s.path
}

func (s *system) Wait() {
	c, cancel := signal.NotifyContext(s.ctx, os.Interrupt, os.Kill)
	defer cancel()
	<-c.Done()
}

func (s *system) Shutdown() {
	s.cancel()
}

func (s *system) SelectActor(path string) Actor {
	return s.actors[path]
}

func (s *system) ActorOf(name string, ref ActorRef) Actor {
	return newActor(s, s.path+"/user", name, ref)
}

func System(parent context.Context, name string) ActorSystem {
	ret := &system{
		name:   name,
		path:   fmt.Sprintf("goka://%s", name),
		actors: make(map[string]Actor),
	}

	ret.ctx, ret.cancel = context.WithCancel(parent)
	return ret
}
