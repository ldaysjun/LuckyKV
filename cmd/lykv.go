package cmd

import (
	"context"
	"fmt"
)

type base struct {
	onClose func()
}

type Client struct {
	base
	commands

	ctx context.Context
}

func NewClient() *Client {
	c := Client{
		base: base{},
		ctx:  context.Background(),
	}

	lykv := 1
	fmt.Println(lykv)


	//c.commands =
	return &c
}
