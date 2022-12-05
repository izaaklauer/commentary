package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/izaaklauer/commentary/config"
	commentaryv1 "github.com/izaaklauer/commentary/gen/proto/go/commentary/v1"
	"github.com/izaaklauer/commentary/server"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("starting commentary.......")
	defer fmt.Println("commentary exiting!")

	err := serve()
	if err != nil {
		log.Fatalf("failed serving\n%s", err)
	}
}

func serve() error {
	var c config.Config

	configPath := os.Getenv("CONFIG_PATH")
	if configPath != "" {
		var err error
		c, err = config.GetConfig(configPath)
		if err != nil {
			return errors.Wrapf(err, "failed to load config from %q", configPath)
		}
	} else {
		// Erroring on no config would be totally valid
		// return errors.New("Environment variable CONFIG_PATH is unset")

		c = config.DefaultConfig()
	}

	// Start the service
	commentaryServer, err := server.NewCommentaryServer(c.Commentary)
	if err != nil {
		return errors.Wrapf(err, "failed to start commentary server")
	}

	listener, err := net.Listen("tcp", c.Server.BindAddr)
	if err != nil {
		return errors.Wrapf(err, "failed to listen on %s", c.Server.BindAddr)
	}
	grpcServer := grpc.NewServer()

	commentaryv1.RegisterCommentaryServiceServer(grpcServer, commentaryServer)
	reflection.Register(grpcServer)

	log.Printf("Serving on %q", c.Server.BindAddr)
	if err := grpcServer.Serve(listener); err != nil {
		return errors.Wrapf(err, "gRPC server exited")
	}

	return nil
}
