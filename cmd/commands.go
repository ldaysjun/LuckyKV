package cmd

import (
	"github.com/ldaysjun/lykv/io"
)

type Cmd interface {
	Err() error
	argsData() []interface{}
	setErr(error)
	readReply(rd *io.Reader) error
}

type CommandsAbler interface {
	Get(key string)
	Set(key string)
}

type baseCmd struct {
	args []interface{}
	err  error
}

func (cmd *baseCmd) Err() error {
	return cmd.err
}

func (cmd *baseCmd) argsData() []interface{} {
	return cmd.args
}

func (cmd *baseCmd) setErr(err error) {
	cmd.err = err
}

func NewStringCmd(args ...interface{}) *StringCmd {
	return &StringCmd{
		baseCmd: baseCmd{
			args: args,
		},
	}
}

type StringCmd struct {
	baseCmd
	val string
}

func (cmd *StringCmd) readReply(rd *io.Reader) error {
	cmd.val, cmd.err = rd.ReadString()
	return cmd.err
}

func (cmd *StringCmd) Result() (string, error) {
	return cmd.val, cmd.err
}

type commands func(cmd Cmd) error

func (c commands) Get(key string) *StringCmd {
	cmd := NewStringCmd("get", key)
	_ = c(cmd)
	return cmd
}

func (c commands) Set(key string, value interface{}) {
	args := make([]interface{}, 3, 3)
	args[0] = "set"
	args[1] = key
	args[2] = value
	cmd := NewStringCmd(args...)
	_ = c(cmd)
}
