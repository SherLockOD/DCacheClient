//Package DCache comment
// This file war generated by tars2go 1.1
// Generated from ProxyShare.tars
package DCache

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//SKeyVersionBS strcut implement
type SKeyVersionBS struct {
	KeyItem []int8 `json:"keyItem"`
	Ver     int8   `json:"ver"`
}

func (st *SKeyVersionBS) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *SKeyVersionBS) ReadFrom(_is *codec.Reader) error {
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
		st.KeyItem = make([]int8, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = _is.Read_int8(&st.KeyItem[i0], 0, false)
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
		err = _is.Read_slice_int8(&st.KeyItem, length, true)
		if err != nil {
			return err
		}

	} else {
		err = fmt.Errorf("require vector, but not")
		if err != nil {
			return err
		}
	}

	err = _is.Read_int8(&st.Ver, 2, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *SKeyVersionBS) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require SKeyVersionBS, but not exist. tag %d", tag)
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
func (st *SKeyVersionBS) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.WriteHead(codec.SIMPLE_LIST, 1)
	if err != nil {
		return err
	}
	err = _os.WriteHead(codec.BYTE, 0)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.KeyItem)), 0)
	if err != nil {
		return err
	}
	err = _os.Write_slice_int8(st.KeyItem)
	if err != nil {
		return err
	}

	err = _os.Write_int8(st.Ver, 2)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *SKeyVersionBS) WriteBlock(_os *codec.Buffer, tag byte) error {
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