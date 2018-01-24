package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	a "github.com/alejandroEsc/grpc-example/api"
	c "github.com/alejandroEsc/grpc-example/configs"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	log.Printf("call to find swagger resource.... %s", r.URL.Path)
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		log.Printf("Not a swagger file %s, missing suffix .swagger.json", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	gwSwaggerDir := c.ParseGWSwaggerEnvVars()

	log.Printf("Serving %s", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(gwSwaggerDir, p)
	http.ServeFile(w, r, p)
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwPort, port, gwAddress, address := c.ParseGateWayEnvVars()
	grpcEndpoint := fmt.Sprintf("%s:%d", address, port)

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", serveSwagger)

	muxGateway := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := a.RegisterHelloServiceHandlerFromEndpoint(ctx, muxGateway, grpcEndpoint, opts)
	if err != nil {
		return err
	}

	mux.Handle("/", muxGateway)

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
		cancel()
		log.Print("service terminated")
		os.Exit(0)
	}()

	return http.ListenAndServe(fmt.Sprintf("%s:%d", gwAddress, gwPort), mux)
}

func main() {
	var err error
	log.Print("starting server")

	err = c.InitEnvVars()
	if err != nil {
		log.Fatalf("failed to init config vars: %s", err)
	}

	if err = run(); err != nil {
		log.Fatal(err)
	}
}
