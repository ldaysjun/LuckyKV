package cmd

type Cmder interface {
	Err() error
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

//-------------------------------------------
type commands func(cmd Cmder) error

func (c commands) Get(key string) {
	cmd := NewStringCmd("get", key)
	_ := c(cmd)
}

//-------------------------------------------
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

//-------------------------------------------
