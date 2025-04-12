package tests

import (
	"context"
	"testing"

	pb "avito-pvz/proto/pvz_v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestGetPVZList(t *testing.T) {
	conn, err := grpc.NewClient("localhost:3000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPVZServiceClient(conn)

	resp, err := client.GetPVZList(context.Background(), &pb.GetPVZListRequest{})
	if err != nil {
		t.Fatalf("Error calling GetPVZList: %v", err)
	}

	t.Log("Response:", resp)
	assert.NotNil(t, resp, "Response should not be nil")
}
