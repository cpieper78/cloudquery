// Code generated by codegen; DO NOT EDIT.

package apikeys

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"

	"cloud.google.com/go/apikeys/apiv2"

	pb "google.golang.org/genproto/googleapis/api/apikeys/v2"

	"google.golang.org/api/option"
)

func createKeys() (*client.Services, error) {
	fakeServer := &fakeKeysServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterApiKeysServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	// Create a client.
	svc, err := apikeys.NewClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &client.Services{
		ApikeysClient: svc,
	}, nil
}

type fakeKeysServer struct {
	pb.UnimplementedApiKeysServer
}

func (f *fakeKeysServer) ListKeys(context.Context, *pb.ListKeysRequest) (*pb.ListKeysResponse, error) {
	resp := pb.ListKeysResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestKeys(t *testing.T) {
	client.MockTestHelper(t, Keys(), createKeys, client.TestOptions{})
}
