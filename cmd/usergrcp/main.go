package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/rostekus/go-graphql/graph/model"
	pb "github.com/rostekus/go-graphql/proto" // Import the generated protobuf package
	"google.golang.org/grpc"
)

var fakeUsers = []model.User{
	{
		ID:       "1",
		Username: "One",
		Password: "OnePass",
	},
	{
		ID:       "2",
		Username: "Two",
		Password: "TwoPass",
	},
}

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.UserResponse, error) {
	log.Printf("Recided %v", in.GetUserId())
	for _, user := range fakeUsers {
		if user.ID == in.GetUserId() {
			return &pb.UserResponse{
				Id:       user.ID,
				Username: user.Username,
				Password: user.Password,
			}, nil
		}
	}
	return &pb.UserResponse{}, fmt.Errorf("no user")
}

func main() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed")
	}
}
