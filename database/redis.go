package database

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rostekus/go-graphql/graph/model"
)

type RedisDB struct {
	client *redis.Client
}

func NewRedisDB(addr, password string, db int) (*RedisDB, error) {
	rdb := &RedisDB{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		}),
	}

	pong, err := rdb.client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to Redis:", pong)

	return rdb, nil
}
func (rdb *RedisDB) SetPost(post *model.Post) error {
	userJSON, err := json.Marshal(post)
	if err != nil {
		return err
	}

	err = rdb.client.Set(context.Background(), "post:"+post.ID, userJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
func (rdb *RedisDB) GetPost(postID string) (*model.Post, error) {
	userJSON, err := rdb.client.Get(context.Background(), "post:"+postID).Result()
	if err != nil {
		return nil, err
	}

	var post model.Post
	err = json.Unmarshal([]byte(userJSON), &post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
