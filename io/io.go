package io

type Reader struct {
	Val interface{}
	Err   error
}

func (r *Reader)ReadString()(string,error)  {
	if r.Err != nil {
		return "",r.Err
	}
	val := r.Val.(string)
	return val,r.Err
}