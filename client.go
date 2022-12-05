package go_trongrid_api

import (
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"google.golang.org/grpc"
)

var defaultNodeAddr = "grpc.shasta.trongrid.io:50051"

type Client struct {
	*client.GrpcClient
}

func New(nodeAddr string, apiKey string) *Client {
	conn := client.NewGrpcClient(nodeAddr)
	if err := conn.Start(grpc.WithInsecure()); err != nil {
		panic(err)
	}
	if len(apiKey) > 0 {
		_ = conn.SetAPIKey(apiKey)
	}

	return &Client{
		conn,
	}
}
