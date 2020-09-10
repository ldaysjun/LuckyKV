package kvql

import (
	"github.com/ldaysjun/lykv/io"
	"github.com/ldaysjun/lykv/storage"
)

type LyString func() error

func (s *LyString) Get(args ...interface{}) *io.Reader {
	kvDB := storage.GetKvDb()
	key := ""
	if len(args) >= 2 {
		key = args[1].(string)
	}
	object, err := kvDB.FindKey(key)
	if err != nil {
		return &io.Reader{
			Val: "",
			Err: err,
		}
	}
	return &io.Reader{
		Val: object.Ptr,
		Err: err,
	}
}

func (s *LyString) Set(args ...interface{}) *io.Reader {
	kvDB := storage.GetKvDb()
	key := args[1].(string)
	val := args[2]
	object := &storage.KvObject{
		Type: storage.LyKvString,
		Ptr:  val,
	}
	reader := new(io.Reader)
	if err := kvDB.SetKey(key, object); err != nil {
		reader.Err = err
		return reader
	}
	reader.Err = nil
	reader.Val = "success"
	return reader
}
