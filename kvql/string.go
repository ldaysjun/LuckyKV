package kvql

import (
	"github.com/ldaysjun/lykv/io"
	"github.com/ldaysjun/lykv/storage"
	"github.com/ldaysjun/lykv/storage/db"
)

type str func() error

func (s *str)Get(args ...interface{}) *io.Reader{
	kvDB := db.GetKvDb()
	key := ""
	if len(args) >= 2 {
		key = args[1].(string)
	}

	object,err := kvDB.FindKey(key)
	if err != nil {
		return &io.Reader{
			Val:"",
			Err:err,
		}
	}

	
	s.parsing(object)
	return nil

}

func (s *str)Set()  {

}

func (s *str)parsing(object *storage.KvObject)  {

}



