package kvql

import (
	"github.com/ldaysjun/lykv/io"
	"github.com/ldaysjun/lykv/storage"
	"github.com/ldaysjun/lykv/storage/db"
)

type str func() error

func (s *str) Get(args ...interface{}) *io.Reader {
	kvDB := db.GetKvDb()
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

func (s *str) Set(args ...interface{}) *io.Reader {
	kvDB := db.GetKvDb()
	key := args[1].(string)
	val := args[2]
	object := &storage.KvObject{
		Type: storage.LYKV_STRING,
		Ptr:  val,
	}
	rd := new(io.Reader)
	if err := kvDB.SetKey(key, object); err != nil {
		rd.Err = err
		return rd
	}
	rd.Err = nil
	rd.Val = "ok"
	return rd
}
