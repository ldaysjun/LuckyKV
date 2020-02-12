package main

import (
	"fmt"
	"github.com/ldaysjun/lykv/cmd"
)

func main() {
	//r := redis.NewClient(nil)
	//r.Get("").Result()
	//r.Set().Result()


	t := cmd.NewClient()
	t.Set("ldjzhang","1")
	s,_ := t.Get("ldjzhang").Result()
	fmt.Println(s)

}

type name struct {
	s string
}
