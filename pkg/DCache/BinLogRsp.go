//Package DCache comment
// This file war generated by tars2go 1.1
// Generated from CacheShare.tars
package DCache

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//BinLogRsp strcut implement
type BinLogRsp struct {
	LogContent []string `json:"logContent"`
	CompLog    string   `json:"compLog"`
	CurLogfile string   `json:"curLogfile"`
	CurSeek    int64    `json:"curSeek"`
	SyncTime   int32    `json:"syncTime"`
	LastTime   int32    `json:"lastTime"`
}

func (st *BinLogRsp) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *BinLogRsp) ReadFrom(_is *codec.Reader) error {
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
		st.LogContent = make([]string, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = _is.Read_string(&st.LogContent[i0], 0, false)
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

	err = _is.Read_string(&st.CompLog, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.CurLogfile, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_int64(&st.CurSeek, 4, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.SyncTime, 5, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.LastTime, 6, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *BinLogRsp) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require BinLogRsp, but not exist. tag %d", tag)
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
func (st *BinLogRsp) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.WriteHead(codec.LIST, 1)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.LogContent)), 0)
	if err != nil {
		return err
	}
	for _, v := range st.LogContent {

		err = _os.Write_string(v, 0)
		if err != nil {
			return err
		}
	}

	err = _os.Write_string(st.CompLog, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.CurLogfile, 3)
	if err != nil {
		return err
	}

	err = _os.Write_int64(st.CurSeek, 4)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.SyncTime, 5)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.LastTime, 6)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *BinLogRsp) WriteBlock(_os *codec.Buffer, tag byte) error {
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
