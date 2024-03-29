//Package DCache comment
// This file war generated by tars2go 1.1
// Generated from ProxyShare.tars
package DCache

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//MKVBatchReq strcut implement
type MKVBatchReq struct {
	ModuleName   string      `json:"moduleName"`
	MainKeys     []string    `json:"mainKeys"`
	Field        string      `json:"field"`
	Cond         []Condition `json:"cond"`
	IdcSpecified string      `json:"idcSpecified"`
}

func (st *MKVBatchReq) resetDefault() {
	st.IdcSpecified = ""
}

//ReadFrom reads  from _is and put into struct.
func (st *MKVBatchReq) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.ModuleName, 1, true)
	if err != nil {
		return err
	}

	err, _, ty = _is.SkipToNoCheck(2, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		st.MainKeys = make([]string, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = _is.Read_string(&st.MainKeys[i0], 0, false)
			if err != nil {
				return err
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

	err = _is.Read_string(&st.Field, 3, true)
	if err != nil {
		return err
	}

	err, _, ty = _is.SkipToNoCheck(4, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		st.Cond = make([]Condition, length, length)
		for i1, e1 := int32(0), length; i1 < e1; i1++ {

			err = st.Cond[i1].ReadBlock(_is, 0, false)
			if err != nil {
				return err
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

	err = _is.Read_string(&st.IdcSpecified, 5, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *MKVBatchReq) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require MKVBatchReq, but not exist. tag %d", tag)
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
func (st *MKVBatchReq) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.ModuleName, 1)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.MainKeys)), 0)
	if err != nil {
		return err
	}
	for _, v := range st.MainKeys {

		err = _os.Write_string(v, 0)
		if err != nil {
			return err
		}
	}

	err = _os.Write_string(st.Field, 3)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 4)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.Cond)), 0)
	if err != nil {
		return err
	}
	for _, v := range st.Cond {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}
	}

	err = _os.Write_string(st.IdcSpecified, 5)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *MKVBatchReq) WriteBlock(_os *codec.Buffer, tag byte) error {
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
