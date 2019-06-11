//Package DCache comment
// This file war generated by tars2go 1.1
// Generated from ProxyShare.tars
package DCache

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//GetMKVRsp strcut implement
type GetMKVRsp struct {
	Data []map[string]string `json:"data"`
}

func (st *GetMKVRsp) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *GetMKVRsp) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err, _, ty = _is.SkipToNoCheck(1, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		st.Data = make([]map[string]string, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err, have = _is.SkipTo(codec.MAP, 0, false)
			if err != nil {
				return err
			}
			if have {
				err = _is.Read_int32(&length, 0, true)
				if err != nil {
					return err
				}
				st.Data[i0] = make(map[string]string)
				for i1, e1 := int32(0), length; i1 < e1; i1++ {
					var k1 string
					var v1 string

					err = _is.Read_string(&k1, 0, false)
					if err != nil {
						return err
					}

					err = _is.Read_string(&v1, 1, false)
					if err != nil {
						return err
					}

					st.Data[i0][k1] = v1
				}
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("not support simple_list type")
		if err != nil {
			return err
		}
	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *GetMKVRsp) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require GetMKVRsp, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *GetMKVRsp) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.WriteHead(codec.LIST, 1)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.Data)), 0)
	if err != nil {
		return err
	}
	for _, v := range st.Data {

		err = _os.WriteHead(codec.MAP, 0)
		if err != nil {
			return err
		}
		err = _os.Write_int32(int32(len(v)), 0)
		if err != nil {
			return err
		}
		for k2, v2 := range v {

			err = _os.Write_string(k2, 0)
			if err != nil {
				return err
			}

			err = _os.Write_string(v2, 1)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//WriteBlock encode struct
func (st *GetMKVRsp) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
