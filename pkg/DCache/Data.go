//Package DCache comment
// This file war generated by tars2go 1.1
// Generated from CacheShare.tars
package DCache

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//Data strcut implement
type Data struct {
	K                SKey   `json:"k"`
	V                SValue `json:"v"`
	ExpireTimeSecond uint32 `json:"expireTimeSecond"`
	Dirty            bool   `json:"dirty"`
	BIsOnlyKey       bool   `json:"bIsOnlyKey"`
}

func (st *Data) resetDefault() {
	st.ExpireTimeSecond = 0
	st.Dirty = true
	st.BIsOnlyKey = false
}

//ReadFrom reads  from _is and put into struct.
func (st *Data) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = st.K.ReadBlock(_is, 1, true)
	if err != nil {
		return err
	}

	err = st.V.ReadBlock(_is, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_uint32(&st.ExpireTimeSecond, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_bool(&st.Dirty, 4, true)
	if err != nil {
		return err
	}

	err = _is.Read_bool(&st.BIsOnlyKey, 5, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *Data) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require Data, but not exist. tag %d", tag)
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
func (st *Data) WriteTo(_os *codec.Buffer) error {
	var err error

	err = st.K.WriteBlock(_os, 1)
	if err != nil {
		return err
	}

	err = st.V.WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	err = _os.Write_uint32(st.ExpireTimeSecond, 3)
	if err != nil {
		return err
	}

	err = _os.Write_bool(st.Dirty, 4)
	if err != nil {
		return err
	}

	err = _os.Write_bool(st.BIsOnlyKey, 5)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *Data) WriteBlock(_os *codec.Buffer, tag byte) error {
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
