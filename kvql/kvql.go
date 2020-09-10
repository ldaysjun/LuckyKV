package kvql

import (
	"errors"

	"github.com/ldaysjun/lykv/io"
)

type KvQL struct {
	LyString
}

// 添加命令
func (kv *KvQL) ParsingCmd(args ...interface{}) *io.Reader {
	if len(args) == 0 {
		return &io.Reader{
			Val: "",
			Err: errors.New("cmd is nil"),
		}
	}
	var reader *io.Reader
	if args[0] == "get" {
		reader = kv.Get(args...)
	}
	if args[0] == "set" {
		reader = kv.Set(args...)
	}
	return reader
}
