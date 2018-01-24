package client

import (
	"log"

	"context"
	"fmt"
	a "github.com/alejandroEsc/grpc-example/api"
	c "github.com/alejandroEsc/grpc-example/configs"
	"google.golang.org/grpc"
)

func runDoKnock(client a.HelloServiceClient) error {
	k := a.Knock{KnockDoor: true}
	log.Print("knocking the door")

	r, err := client.GetHello(context.Background(), &k)

	if err != nil {
		return err
	}

	log.Printf("knocked the door, got a reply: %s", r.ReplyMessage)
	return nil
}

func runDoNotKnock(client a.HelloServiceClient) error {
	k := a.Knock{KnockDoor: false}
	log.Print("NOT knocking the door")

	r, err := client.GetHello(context.Background(), &k)
	if err != nil {
		return err
	}

	log.Printf("Standing in front of the door, got a message?: %s", r.ReplyMessage)
	return nil
}

func runNoMessage(client a.HelloServiceClient) error {
	log.Print("sending nil message")

	r, err := client.GetHello(context.Background(), nil)
	if err != nil {
		return err
	}

	log.Printf("no message sent, and got a reply: %s", r.ReplyMessage)
	return nil

}

func main() {
	if err := c.InitEnvVars(); err != nil {
		log.Fatalf("failed to init config vars: %s", err)
	}

	port, _, address := c.ParseEnvVars()

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", address, port), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := a.NewHelloServiceClient(conn)

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
