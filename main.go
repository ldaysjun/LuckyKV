package main

import (
	"github.com/go-redis/redis/v7"
	"github.com/ldaysjun/lykv/cmd"
	"github.com/ldaysjun/lykv/kvql"
)

func main() {
	r := redis.NewClient(nil)
	r.Get("").Result()



	t := cmd.NewClient()
	t.Get("test")


	k := kvql.KvQL{}
}

type name struct {
	s string
}
