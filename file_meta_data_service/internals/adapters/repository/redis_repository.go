package repository

import (
	"context"
	"filemetadata-service/internals/core/ports"
	"os"

	"github.com/redis/go-redis/v9"
)

type RedisDBRepository struct {
	Client *redis.Client
	Ctx    context.Context
	Key    string

	// This implementation should be
	// removed when RedisDBRepository
	// Implements all method of FileMetaDataDBRepository
	ports.FileMetaDataDBRepository
}

func NewRedisDBRepository() *RedisDBRepository {

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	return &RedisDBRepository{Client: client, Ctx: ctx}
}

// this function can serve as both the key
// to the value needed to be stored
func (rDb *RedisDBRepository) Table(key string) ports.FileMetaDataDBRepository {

	rDb.Key = key

	return rDb
}

func (rDb *RedisDBRepository) Get() (any, error) {

	val, err := rDb.Client.Get(rDb.Ctx, rDb.Key).Result()

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (rDb *RedisDBRepository) Create(data any) (any, error) {

	err := rDb.Client.Set(rDb.Ctx, rDb.Key, data, 0).Err()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (rDb *RedisDBRepository) Delete() error {
	rDb.Client.Del(rDb.Ctx, rDb.Key).Result()
	return nil
}
