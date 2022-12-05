package server

import (
	"context"
	"log"

        "github.com/izaaklauer/commentary/config"
	commentaryv1 "github.com/izaaklauer/commentary/gen/proto/go/commentary/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CommentaryServer struct {
	commentaryv1.UnimplementedCommentaryServiceServer

	config config.Commentary
}

// NewCommentaryServer initializes a new server from config
func NewCommentaryServer(config config.Commentary) (*CommentaryServer, error) {
	// Server-specific initialization, like DB clients, goes here.

	server := CommentaryServer{
		config: config,
	}

	return &server, nil
}

func (s * CommentaryServer) HelloWorld(
	ctx context.Context,
	req *commentaryv1.HelloWorldRequest,
) (*commentaryv1.HelloWorldResponse, error) {
	log.Printf("HelloWorld request with message %q", req.Message)

	resp := &commentaryv1.HelloWorldResponse{
		RequestMessage: req.Message,
		ConfigMessage:  s.config.HelloWorldMessage,
		Now:            timestamppb.Now(),
	}

	return resp, nil
}
