package service

import (
	"fmt"

	"github.com/rostekus/go-graphql/database"
	"github.com/rostekus/go-graphql/graph/model"
	// Import the generated protobuf package
)

type IUserService interface {
	Get(id string) (*model.User, error)
}

type UserService struct {
	Db database.IDatabase
}

func (src *UserService) Get(id string) (*model.User, error) {
	usr, err := src.Db.GetUser(id)

	if err != nil {
		return nil, fmt.Errorf("couldn't retrieve from db: %v", err)

	}
	return usr, nil

}
