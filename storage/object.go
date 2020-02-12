package storage

type KvObject struct {
	Type     KvType
	Encoding string
	Ptr      interface{}
}


