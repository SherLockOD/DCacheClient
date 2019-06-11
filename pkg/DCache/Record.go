//Package DCache comment
// This file war generated by tars2go 1.1
// Generated from ProxyShare.tars
package DCache

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//Record strcut implement
type Record struct {
	MainKey  string            `json:"mainKey"`
	MpRecord map[string]string `json:"mpRecord"`
	Ret      int32             `json:"ret"`
}

func (st *Record) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *Record) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.MainKey, 1, true)
	if err != nil {
		return err
	}

	err, have = _is.SkipTo(codec.MAP, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&length, 0, true)
	if err != nil {
		return err
	}
	st.MpRecord = make(map[string]string)
	for i0, e0 := int32(0), length; i0 < e0; i0++ {
		var k0 string
		var v0 string

		err = _is.Read_string(&k0, 0, false)
		if err != nil {
			return err
		}

		err = _is.Read_string(&v0, 1, false)
		if err != nil {
			return err
		}

		st.MpRecord[k0] = v0
	}

	err = _is.Read_int32(&st.Ret, 3, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *Record) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require Record, but not exist. tag %d", tag)
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
func (st *Record) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.MainKey, 1)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.MAP, 2)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.MpRecord)), 0)
	if err != nil {
		return err
	}
	for k1, v1 := range st.MpRecord {

		err = _os.Write_string(k1, 0)
		if err != nil {
			return err
		}

		err = _os.Write_string(v1, 1)
		if err != nil {
			return err
		}
	}

	err = _os.Write_int32(st.Ret, 3)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *Record) WriteBlock(_os *codec.Buffer, tag byte) error {
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
