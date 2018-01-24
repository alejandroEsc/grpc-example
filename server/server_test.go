package main

import (
	"context"
	"testing"

	a "github.com/alejandroEsc/grpc-example/api"
)

var (
	noKnockMsg = "try and knock the door"
	repMsg     = "Hello!"
)

func TestGetHello(t *testing.T) {
	ds := doorServer{knockFailureMsg: noKnockMsg}

	req := &a.Knock{KnockDoor: true}

	resp, err := ds.GetHello(context.Background(), req)

	if err != nil {
		t.Errorf("got an unexpected error: %s", err)
	}

	if resp.ReplyMessage != repMsg {
		t.Errorf("got an unexpected reply: %s %s %s", resp.ReplyMessage, "intead of ", repMsg)
	}
}
