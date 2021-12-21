package goka_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"context"

	"github.com/dairaga/goka"
)

type Interest string
type Three string
type Two string
type Why string
type Because string

type Penguin struct{}

func (p *Penguin) Receive(ctx goka.Context, msg interface{}) {
	fmt.Println(ctx.Sender().Name(), "->", ctx.Name(), ":", msg)
	switch msg.(type) {
	case Interest:
		ctx.Reply(Three("eat, sleep, beat dongdong."))
	}
}

type DongDong struct{}

func (p *DongDong) Receive(ctx goka.Context, msg interface{}) {
	fmt.Println(ctx.Sender().Name(), "->", ctx.Name(), ":", msg)
	switch msg.(type) {
	case Interest:
		ctx.Reply(Two("eat, sleep."))
	case Why:
		ctx.Reply(Because("I'm DongDong."))
	}
}

type Reporter struct{}

func (r *Reporter) Receive(ctx goka.Context, msg interface{}) {
	fmt.Println(ctx.Sender().Name(), "->", ctx.Name(), ":", msg)
	switch msg.(type) {
	case Three:
	case Two:
		ctx.Reply(Why("why"))
	case Because:
	}
}

func Example() {
	system := goka.System(context.Background(), "mytest")
	fmt.Println(system.Name())
	fmt.Println(system.Path())

	// Output:
	// mytest
	// goka://mytest
}

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	system := goka.System(ctx, "mytest")
	fmt.Println(system.Name())
	fmt.Println(system.Path())

	p := new(Penguin)
	d := new(DongDong)
	r := new(Reporter)

	report := system.ActorOf("reporter", r)
	report.Send(system.ActorOf("penguin", p), Interest("interest"))
	report.Send(system.ActorOf("dongdong", d), Interest("interest"))

	system.Wait()
	system.Shutdown()
	os.Exit(m.Run())
}
