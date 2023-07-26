package database

import "github.com/rostekus/go-graphql/graph/model"

type IDatabase interface {
	GetUser(string) (*model.User, error)
}
