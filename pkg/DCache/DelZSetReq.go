//Package DCache comment
// This file war generated by tars2go 1.1
// Generated from ProxyShare.tars
package DCache

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//DelZSetReq strcut implement
type DelZSetReq struct {
	ModuleName string      `json:"moduleName"`
	MainKey    string      `json:"mainKey"`
	Cond       []Condition `json:"cond"`
}

func (st *DelZSetReq) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *DelZSetReq) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.ModuleName, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.MainKey, 2, true)
	if err != nil {
		return err
	}

	err, _, ty = _is.SkipToNoCheck(3, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		st.Cond = make([]Condition, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = st.Cond[i0].ReadBlock(_is, 0, false)
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

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *DelZSetReq) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require DelZSetReq, but not exist. tag %d", tag)
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
func (st *DelZSetReq) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.ModuleName, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.MainKey, 2)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 3)
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

	return nil
}

//WriteBlock encode struct
func (st *DelZSetReq) WriteBlock(_os *codec.Buffer, tag byte) error {
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