package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/ankit/project/grpc/grpc-server-streaming/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAppServiceServer
}

func isPrime(no int32) bool {
	if no < 2 {
		return false
	}
	for i := int32(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func (s *server) GeneratePrimes(req *proto.PrimeRequest, stream proto.AppService_GeneratePrimesServer) error {

	start := req.GetStart()
	end := req.GetEnd()
	fmt.Printf("Generating primes from %d and %d\n", start, end)
	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Printf("Sending Prime No : %d\n", i)
			res := &proto.PrimeResponse{
				PrimeNo: i,
			}
			err := stream.Send(res)
			if err != nil {
				return err
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func main() {
	s := &server{}
	listner, err := net.Listen("tcp", ":50501")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, s)
	grpcServer.Serve(listner)

}
