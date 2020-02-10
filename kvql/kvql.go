package kvql

import (
	"errors"
	"github.com/ldaysjun/lykv/io"
)

type KvQL struct {
	str
}

func (kv *KvQL) ParsingCmd(args ...interface{}) *io.Reader {
	if len(args) == 0 {
		return &io.Reader{
			Val: "",
			Err: errors.New("cmd is nil"),
		}
	}
	if args[0] == "get" {
		kv.Get(args)
	}
	return &io.Reader{}
}
