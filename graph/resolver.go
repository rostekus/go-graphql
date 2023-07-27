package graph

import "github.com/rostekus/go-graphql/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService service.IUserService
	PostService service.IPostService
}
