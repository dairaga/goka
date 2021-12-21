package main

import (
	"fmt"

	"github.com/dairaga/goka"
)

type report struct{}

func (r *report) Receive(sender goka.Actor, data any) {
	switch v := data.(type) {
	case string:
		fmt.Println(v)
	}
}

type penguin struct{}

func (p *penguin) Receive(sender goka.Actor, data any) {
	switch v := data.(type) {
	case string:
		fmt.Println(v)
		sender.
	}
}

func main() {
	system := goka.System("test")

}
