package main

import (
	"context"
	"io"
	"log"

	"github.com/ankit/project/grpc/grpc-server-streaming/proto"
	"google.golang.org/grpc"
)

func doServerStreaming(ctx context.Context, client proto.AppServiceClient) {
	req := &proto.PrimeRequest{
		Start: 6,
		End:   99,
	}
	stream, err := client.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Prime No = ", res.GetPrimeNo())
	}

}

func main() {
	clientConn, err := grpc.Dial("localhost:50501", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	//proxy instance
	client := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	//server Streaming
	doServerStreaming(ctx, client)
}
