package slick

import (
	"testing"
	"time"

	"github.com/nlopes/slack"
)

func TestListenerCheckParams(t *testing.T) {
	c := Listener{
		ListenUntil:    time.Now(),
		ListenDuration: 120 * time.Second,
	}
	err := c.checkParams()
	if err == nil {
		t.Error("checkParams shouldn't be nil")
	}
}

func TestDefeaultFilter(t *testing.T) {
	c := &Listener{}
	u := &slack.User{ID: "a_user"}
	m := &Message{Msg: &slack.Msg{Text: "hello mama"}, FromUser: u}

	if c.filterMessage(m) != true {
		t.Error("filterMessage Failed")
	}

	type El struct {
		c *Listener
		r bool
	}
	tests := []El{
		El{&Listener{}, true},

		El{&Listener{
			Contains: "moo",
		}, false},

		El{&Listener{
			Contains: "MAMA",
		}, true},

		El{&Listener{
			WithUser: u,
		}, true},

		El{&Listener{
			WithUser: &slack.User{ID: "another_user"},
		}, false},
	}

	for i, el := range tests {
		if el.c.filterMessage(m) != el.r {
			t.Error("filterMessage Failed, index ", i)
		}
	}
}
