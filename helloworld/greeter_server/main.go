package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/vglvishnu/gosamples/helloworld/helloworldproto"
)
const (
	port = ":50051"
)
type server struct{
	pb.UnimplementedGreetServiceServer
}
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received greet from :: %+v",in.GetName())
	return &pb.HelloReply{Message: "Hello "+ in.GetName()}, nil
}
func main(){
	lis,err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("Failed to listen to port :: %+v\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s,&server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %+v", err)
	}
}