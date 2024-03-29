//Package DCache comment
// This file war generated by tars2go 1.1
// Generated from CacheShare.tars
package DCache

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//DbUpdateValue strcut implement
type DbUpdateValue struct {
	Op    Op       `json:"op"`
	Value string   `json:"value"`
	Type  DataType `json:"type"`
}

func (st *DbUpdateValue) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *DbUpdateValue) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_int32((*int32)(&st.Op), 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Value, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32((*int32)(&st.Type), 3, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *DbUpdateValue) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require DbUpdateValue, but not exist. tag %d", tag)
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
func (st *DbUpdateValue) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(int32(st.Op), 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Value, 2)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(st.Type), 3)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *DbUpdateValue) WriteBlock(_os *codec.Buffer, tag byte) error {
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
