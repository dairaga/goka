package goka

import (
	"context"
	"fmt"
)

type ActorSystem interface {
	Name() string
	Path() string
	ActorOf(name string, ref ActorRef) Actor
	Select(path string) Actor
	Wait()
	Shutdown()
}

type system struct {
	_ struct{}

	actors map[string]Actor

	name string
	path string

	ctx context.Context
}

var _ ActorSystem = &system{}

func (s *system) Name() string {
	return s.name
}

func (s *system) Path() string {
	return s.path
}

func (s *system) ActorOf(name string, ref ActorRef) Actor {
	return nil
}

func (s *system) Select(path string) Actor {
	return s.actors[path]
}

func (s *system) Wait() {
	select {
	case <-s.ctx.Done():
	}
}

func (s *system) Shutdown() {

}

func System(ctx context.Context, name string) ActorSystem {
	ret := &system{
		actors: make(map[string]Actor),
		name:   name,
		path:   fmt.Sprintf(`goka://%s`, name),
	}
	ret.ctx = context.WithValue(ctx, keySystem, ret)
	return ret
}
