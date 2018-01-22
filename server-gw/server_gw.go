package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
    "os"
    "os/signal"
	"path"
	"strings"
    "syscall"
    "time"

	gw "github.com/alejandroEsc/grpc-example/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

var (
	port         = flag.Int("port", 8502, "Server port.")
	grpcEndpoint = flag.String("echo_endpoint", "localhost:8501", "endpoint of Hello Service")
	swaggerDir = flag.String("swagger-dir", "swagger", "path to the directory which contains swagger definitions")

)

type doorServer struct {
	knockFailureMsg string
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	log.Printf("call to find swagger resource.... %s", r.URL.Path)
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		log.Printf("Not a swagger file %s, missing suffix .swagger.json", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	log.Printf("Serving %s", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(*swaggerDir, p)
	http.ServeFile(w, r, p)
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", serveSwagger)


	muxGateway := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterHelloServiceHandlerFromEndpoint(ctx, muxGateway, *grpcEndpoint, opts)
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
		time.Sleep(2*time.Second)
		cancel()
		log.Print("service terminated")
		os.Exit(0)
	}()



	return http.ListenAndServe(fmt.Sprintf("localhost:%d", *port), mux)
}

func main() {
    var err error
    log.Print("starting server")

    // Parse arguments here
	flag.Parse()

	if err = run(); err != nil {
		log.Fatal(err)
	}
}


