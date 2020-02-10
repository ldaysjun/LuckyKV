package cmd

import "github.com/ldaysjun/lykv/io"

type Cmder interface {
	Err() error
	Args() []interface{}
	setErr(error)
	readReply(rd *io.Reader) error
}

type CommandsAbler interface {
	Get(key string)
	ZRange(key string)
}

type baseCmd struct {
	args []interface{}
	err  error
}

func (cmd *baseCmd) Err() error {
	return cmd.err
}

func (cmd *baseCmd) Args() []interface{} {
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

func (cmd *StringCmd)readReply(rd *io.Reader) error{
	cmd.val,cmd.err = rd.ReadString()
	return cmd.err
}

type commands func(cmd Cmder) error

func (c commands) Get(key string) {
	cmd := NewStringCmd("get", key)
	c(cmd)
}

func (c commands) ZRange(key string) {
	cmd := NewStringCmd("get", key)
	c(cmd)
}



