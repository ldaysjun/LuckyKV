package cmd

import (
	"context"
	"errors"
)

type baseClient struct {
	onClose func()
}

func (base *baseClient) process(ctx context.Context, cmd Cmder) {
	if err := base._process(ctx, cmd); err != nil {
		cmd.SetErr(err)
	}
}

func (base *baseClient) _process(ctx context.Context, cmd Cmder) error {

	return nil
}

func NewClient() *Client {
	c := Client{
		baseClient: baseClient{},
		ctx:        context.Background(),
	}
	c.commands = c.Process
	return &c
}

type Client struct {
	baseClient
	commands

	ctx context.Context
}

func (c *Client)Process(cmd Cmder) error  {
	c.baseClient.process(c.ctx,cmd)
	return errors.New("试试")
}





