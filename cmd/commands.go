package cmd

type Cmder interface {
	Err() error
	SetErr(error)
}

type CommandsAbler interface {
	Get(key string)
}

type baseCmd struct {
	args []interface{}
	err  error
}

func (cmd *baseCmd) Err() error {
	return cmd.err
}


func (cmd *baseCmd) SetErr(err error) {
	cmd.err = err
}

type StringCmd struct {
	baseCmd
	val string
}

func NewStringCmd(args ...interface{}) *StringCmd {
	return &StringCmd{
		baseCmd: baseCmd{
			args: args,
		},
	}
}


type commands func(cmd Cmder) error

func (c commands) Get(key string) {
	cmd := NewStringCmd("get", key)
	c(cmd)
}



