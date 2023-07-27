package service

import (
	"fmt"

	"github.com/rostekus/go-graphql/database"
	"github.com/rostekus/go-graphql/graph/model"
)

type IPostService interface {
	Get(id string) (*model.Post, error)
}

type PostService struct {
	Db database.IPostDatabase
}

func (src *PostService) Get(id string) (*model.Post, error) {
	usr, err := src.Db.GetPost(id)

	if err != nil {
		return nil, fmt.Errorf("couldn't retrieve from db: %v", err)

	}
	return usr, nil

}
