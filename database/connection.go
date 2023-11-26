package database

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var DBClient *redis.Client

func Connect(){
    dsn := os.Getenv("DSN")
    opt, err := redis.ParseURL(dsn)
if err != nil {
	panic(err)
}

    DBClient = redis.NewClient(opt)
}
