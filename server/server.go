package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	a "github.com/alejandroEsc/grpc-example/api"
	c "github.com/alejandroEsc/grpc-example/configs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type doorServer struct {
	knockFailureMsg string
}

// Implemented the server interface
func (s *doorServer) GetHello(c context.Context, knock *a.Knock) (*a.Reply, error) {

	if knock == nil {
		return nil, fmt.Errorf("nothing received, wont respond")
	}

	r := a.Reply{Reply: false, ReplyMessage: s.knockFailureMsg}

	if knock.KnockDoor {
		r.Reply = true
		r.ReplyMessage = "Hello!"
	}

	return &r, nil
}

func newDoorServer(noKnockMsg string) *doorServer {
	d := doorServer{knockFailureMsg: noKnockMsg}
	return &d
}

func main() {
	var err error
	log.Print("starting server")

	err = c.InitEnvVars()
	if err != nil {
		log.Fatalf("failed to init config vars: %s", err)
	}

	port, knockFailure, address := c.ParseEnvVars()

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	a.RegisterHelloServiceServer(grpcServer, newDoorServer(knockFailure))

	log.Printf("attempting to start server in port %d", port)

	//  Get notified that server is being asked to stop
	// Handle SIGINT and SIGTERM.
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	// Chance here to gracefully handle being stopped.
	go func() {
		sig := <-gracefulStop
		log.Printf("caught sig: %+v", sig)
		log.Println("Wait for 2 second to finish processing")
		time.Sleep(2 * time.Second)
		grpcServer.Stop()
		log.Print("service terminated")
		os.Exit(0)
	}()

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("could not start service: %s", err)
	}
}
