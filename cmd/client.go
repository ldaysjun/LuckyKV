package cmd

import "context"

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

	//c.commands =
	return &c
}
