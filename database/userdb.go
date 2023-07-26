package database

import (
	"context"

	"github.com/rostekus/go-graphql/graph/model"
	pb "github.com/rostekus/go-graphql/proto"
	"google.golang.org/grpc"
)

type UserDB struct {
	connString string
}

func NewUserDB(connString string) *UserDB {
	return &UserDB{
		connString: connString,
	}
}

func (db *UserDB) getConn() (*grpc.ClientConn, func() error, error) {
	conn, err := grpc.Dial(db.connString, grpc.WithInsecure())
	return conn, conn.Close, err
}

func (db *UserDB) GetUser(id string) (*model.User, error) {
	conn, closeFun, err := db.getConn()
	defer closeFun()
	client := pb.NewUserServiceClient(conn)

	req := &pb.GetUserRequest{UserId: id}

	user, err := client.GetUser(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return &model.User{
		Password: user.Password,
		Username: user.Username,
	}, nil
}
