package main

import (
	"fmt"

	"github.com/ldaysjun/lykv/cmd"
)

func main() {
	t := cmd.NewClient()
	t.Set("ldjzhang", "1")
	s, err := t.Get("ldjzhang").Result()

	fmt.Println(s)
	fmt.Println(err)
}

type name struct {
	s string
}
