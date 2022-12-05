package go_trongrid_api

import (
	"google.golang.org/grpc"
	"os"
	"testing"
)

var conn *Client

func TestMain(m *testing.M) {
	conn = New(defaultNodeAddr, "")
	if err := conn.Start(grpc.WithInsecure()); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
