//Package DCache comment
// This file war generated by tars2go 1.1
// Generated from ProxyShare.tars
package DCache

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//MKVBatchWriteRsp strcut implement
type MKVBatchWriteRsp struct {
	RspData map[int32]int32 `json:"rspData"`
}

func (st *MKVBatchWriteRsp) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *MKVBatchWriteRsp) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err, have = _is.SkipTo(codec.MAP, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&length, 0, true)
	if err != nil {
		return err
	}
	st.RspData = make(map[int32]int32)
	for i0, e0 := int32(0), length; i0 < e0; i0++ {
		var k0 int32
		var v0 int32

		err = _is.Read_int32(&k0, 0, false)
		if err != nil {
			return err
		}

		err = _is.Read_int32(&v0, 1, false)
		if err != nil {
			return err
		}

		st.RspData[k0] = v0
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *MKVBatchWriteRsp) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require MKVBatchWriteRsp, but not exist. tag %d", tag)
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
func (st *MKVBatchWriteRsp) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.WriteHead(codec.MAP, 1)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.RspData)), 0)
	if err != nil {
		return err
	}
	for k1, v1 := range st.RspData {

		err = _os.Write_int32(k1, 0)
		if err != nil {
			return err
		}

		err = _os.Write_int32(v1, 1)
		if err != nil {
			return err
		}
	}

	return nil
}

//WriteBlock encode struct
func (st *MKVBatchWriteRsp) WriteBlock(_os *codec.Buffer, tag byte) error {
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