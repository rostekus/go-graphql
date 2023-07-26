package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/rostekus/go-graphql/graph/model"
	pb "github.com/rostekus/go-graphql/proto" // Import the generated protobuf package
	"google.golang.org/grpc"
)

func GetUser() *model.User {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	userID := "your_user_id_here"
	req := &pb.GetUserRequest{UserId: userID}

	user, err := client.GetUser(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	fmt.Printf("User ID: %s\n", user.Id)
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Password: %s\n", user.Password)
	return &model.User{
		Password: user.Password,
		Username: user.Username,
	}
}
