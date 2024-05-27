/**
 * @author:伯约
 * @date:2024/5/27
 * @note:
**/

package server

import (
	"context"
	pb "github.com/webxiaohua/GrpcDebug/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewGrpcServer() (s *grpc.Server) {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s = grpc.NewServer()
	pb.RegisterGreeterServer(s, &DemoServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return s
}

type DemoServer struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *DemoServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	resp := &pb.HelloReply{Message: "Hello " + in.GetName()}
	resp.BookList = make([]*pb.Book, 0)
	resp.BookList = append(resp.BookList, &pb.Book{Id: 1, Name: "Head first"})
	return resp, nil
}
