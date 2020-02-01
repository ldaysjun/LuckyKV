package main

import (
	"github.com/go-redis/redis/v7"
	"github.com/ldaysjun/lykv/cmd"
)

func main() {
	r := redis.NewClient(nil)
	r.Get()
	r.Set()
	t := cmd.NewClient()
	t.Get("test")

}