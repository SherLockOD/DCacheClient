//Package DCache comment
// This file war generated by tars2go 1.1
// Generated from ProxyShare.tars
package DCache

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//UpdateKeyValueBS strcut implement
type UpdateKeyValueBS struct {
	MainKey          []int8                       `json:"mainKey"`
	MpValue          map[string]UpdateFieldInfoBS `json:"mpValue"`
	Ver              int8                         `json:"ver"`
	Dirty            bool                         `json:"dirty"`
	Insert           bool                         `json:"insert"`
	ExpireTimeSecond int32                        `json:"expireTimeSecond"`
}

func (st *UpdateKeyValueBS) resetDefault() {
	st.Ver = 0
	st.Dirty = true
	st.Insert = false
	st.ExpireTimeSecond = 0
}

//ReadFrom reads  from _is and put into struct.
func (st *UpdateKeyValueBS) ReadFrom(_is *codec.Reader) error {
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
		st.MainKey = make([]int8, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = _is.Read_int8(&st.MainKey[i0], 0, false)
			if err != nil {
				return err
			}
		}
	} else if ty == codec.SIMPLE_LIST {

		err, _ = _is.SkipTo(codec.BYTE, 0, true)
		if err != nil {
			return err
		}
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		err = _is.Read_slice_int8(&st.MainKey, length, true)
		if err != nil {
			return err
		}

	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return err
		}
	}

	err, have = _is.SkipTo(codec.MAP, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&length, 0, true)
	if err != nil {
		return err
	}
	st.MpValue = make(map[string]UpdateFieldInfoBS)
	for i1, e1 := int32(0), length; i1 < e1; i1++ {
		var k1 string
		var v1 UpdateFieldInfoBS

		err = _is.Read_string(&k1, 0, false)
		if err != nil {
			return err
		}

		err = v1.ReadBlock(_is, 1, false)
		if err != nil {
			return err
		}

		st.MpValue[k1] = v1
	}

	err = _is.Read_int8(&st.Ver, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_bool(&st.Dirty, 4, true)
	if err != nil {
		return err
	}

	err = _is.Read_bool(&st.Insert, 5, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.ExpireTimeSecond, 6, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *UpdateKeyValueBS) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require UpdateKeyValueBS, but not exist. tag %d", tag)
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
func (st *UpdateKeyValueBS) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.WriteHead(codec.SIMPLE_LIST, 1)
	if err != nil {
		return err
	}
	err = _os.WriteHead(codec.BYTE, 0)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.MainKey)), 0)
	if err != nil {
		return err
	}
	err = _os.Write_slice_int8(st.MainKey)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.MAP, 2)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.MpValue)), 0)
	if err != nil {
		return err
	}
	for k2, v2 := range st.MpValue {

		err = _os.Write_string(k2, 0)
		if err != nil {
			return err
		}

		err = v2.WriteBlock(_os, 1)
		if err != nil {
			return err
		}
	}

	err = _os.Write_int8(st.Ver, 3)
	if err != nil {
		return err
	}

	err = _os.Write_bool(st.Dirty, 4)
	if err != nil {
		return err
	}

	err = _os.Write_bool(st.Insert, 5)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.ExpireTimeSecond, 6)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *UpdateKeyValueBS) WriteBlock(_os *codec.Buffer, tag byte) error {
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
