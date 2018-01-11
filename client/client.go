package main

import (
	"context"
	"flag"
	"log"

	a "github.com/alejandroEsc/grpc-example/api"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String(
		"server_addr",
		"127.0.0.1:8501",
		"The server address in the format of host:port",
	)
)

func runDoKnock(client a.HelloClient) error {
	k := a.Knock{Knocked: true}
	log.Print("knocking the door")

	r, err := client.Knocked(context.Background(), &k)

	if err != nil {
		return err
	}

	log.Printf("knocked the door, got a reply: %s", r.Message)
	return nil
}

func runDoNotKnock(client a.HelloClient) error {
	k := a.Knock{Knocked: false}
	log.Print("NOT knocking the door")

	r, err := client.Knocked(context.Background(), &k)
	if err != nil {
		return err
	}

	log.Printf("Standing in front of the door, got a message?: %s", r.Message)
	return nil
}

func runNoMessage(client a.HelloClient) error {
	log.Print("sending nil message")

	r, err := client.Knocked(context.Background(), nil)
	if err != nil {
		return err
	}

	log.Printf("no message sent, and got a reply: %s", r.Message)
	return nil

}

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer func() {
		e := conn.Close()
		if e != nil {
			log.Fatal(e)
		}
	}()

	client := a.NewHelloClient(conn)

	err = runDoKnock(client)
	if err != nil {
		log.Printf("got an error message: %s", err)
	}

	err = runDoNotKnock(client)
	if err != nil {
		log.Printf("got an error message: %s", err)
	}

	err = runNoMessage(client)
	if err != nil {
		log.Printf("got an error message: %s", err)
	}
}
