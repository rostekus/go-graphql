package database

import "github.com/rostekus/go-graphql/graph/model"

type IUserDatabase interface {
	GetUser(string) (*model.User, error)
}

type IPostDatabase interface {
	GetPost(string) (*model.Post, error)
}
