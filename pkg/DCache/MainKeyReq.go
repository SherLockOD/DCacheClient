//Package DCache comment
// This file war generated by tars2go 1.1
// Generated from ProxyShare.tars
package DCache

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//MainKeyReq strcut implement
type MainKeyReq struct {
	ModuleName   string `json:"moduleName"`
	MainKey      string `json:"mainKey"`
	IdcSpecified string `json:"idcSpecified"`
}

func (st *MainKeyReq) resetDefault() {
	st.IdcSpecified = ""
}

//ReadFrom reads  from _is and put into struct.
func (st *MainKeyReq) ReadFrom(_is *codec.Reader) error {
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

	err = _is.Read_string(&st.IdcSpecified, 3, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *MainKeyReq) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require MainKeyReq, but not exist. tag %d", tag)
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
func (st *MainKeyReq) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.ModuleName, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.MainKey, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.IdcSpecified, 3)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *MainKeyReq) WriteBlock(_os *codec.Buffer, tag byte) error {
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