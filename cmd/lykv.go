package cmd

import (
	"context"
	"errors"
	"github.com/ldaysjun/lykv/io"
	"github.com/ldaysjun/lykv/kvql"
)

type baseClient struct {
	onClose func()
	kvQL    kvql.KvQL
}

func (base *baseClient) process(ctx context.Context, cmd Cmder) {
	if err := base._process(ctx, cmd); err != nil {
		cmd.setErr(err)
	}
}

func (base *baseClient) _process(ctx context.Context, cmd Cmder) error {
	args := cmd.argsData()
	rd := base.kvQL.ParsingCmd(args...)
	err := base.withReader(ctx,rd,cmd.readReply)
	if err != nil {
		return err
	}
	return nil
}

func (base *baseClient) withReader(ctx context.Context, rd *io.Reader, fn func(reader *io.Reader) error) error{
	return fn(rd)
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

func (c *Client) Process(cmd Cmder) error {
	c.baseClient.process(c.ctx, cmd)
	return errors.New("试试")
}
