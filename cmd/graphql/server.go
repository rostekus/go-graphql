package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rostekus/go-graphql/database"
	"github.com/rostekus/go-graphql/graph"
	"github.com/rostekus/go-graphql/service"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := database.NewUserDB("localhost:5000")
	userService := service.UserService{Db: db}

	redisDB, err := database.NewRedisDB("localhost:6379", "mysecurepassword", 0)
	if err != nil {
		log.Fatal("cannot connect to redis")
	}
	postService := service.PostService{
		Db: redisDB,
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{UserService: &userService, PostService: &postService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
