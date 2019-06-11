//Package DCache comment
// This file war generated by tars2go 1.1
// Generated from BackUp.tars
package DCache

import (
	"context"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
)

//BackUp struct
type BackUp struct {
	s m.Servant
}

//Dump is the proxy function for the method defined in the tars file, with the context
func (_obj *BackUp) Dump(DumpPath string, MirrorName string, Errmsg *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(DumpPath, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string(MirrorName, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "dump", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*Errmsg), 3, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//DumpWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *BackUp) DumpWithContext(ctx context.Context, DumpPath string, MirrorName string, Errmsg *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(DumpPath, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_string(MirrorName, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "dump", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*Errmsg), 3, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Restore is the proxy function for the method defined in the tars file, with the context
func (_obj *BackUp) Restore(MirrorPath string, BinlogPath []string, LastTime int32, Normal bool, Errmsg *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(MirrorPath, 1)
	if err != nil {
		return ret, err
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return ret, err
	}
	err = _os.Write_int32(int32(len(BinlogPath)), 0)
	if err != nil {
		return ret, err
	}
	for _, v := range BinlogPath {

		err = _os.Write_string(v, 0)
		if err != nil {
			return ret, err
		}
	}

	err = _os.Write_int32(LastTime, 3)
	if err != nil {
		return ret, err
	}

	err = _os.Write_bool(Normal, 4)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "restore", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*Errmsg), 5, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//RestoreWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *BackUp) RestoreWithContext(ctx context.Context, MirrorPath string, BinlogPath []string, LastTime int32, Normal bool, Errmsg *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(MirrorPath, 1)
	if err != nil {
		return ret, err
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return ret, err
	}
	err = _os.Write_int32(int32(len(BinlogPath)), 0)
	if err != nil {
		return ret, err
	}
	for _, v := range BinlogPath {

		err = _os.Write_string(v, 0)
		if err != nil {
			return ret, err
		}
	}

	err = _os.Write_int32(LastTime, 3)
	if err != nil {
		return ret, err
	}

	err = _os.Write_bool(Normal, 4)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "restore", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*Errmsg), 5, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SetServant sets servant for the service.
func (_obj *BackUp) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *BackUp) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}
func (_obj *BackUp) setMap(l int, res *requestf.ResponsePacket, ctx map[string]string, sts map[string]string) {
	if l == 1 {
		for k, _ := range ctx {
			delete(ctx, k)
		}
		for k, v := range res.Context {
			ctx[k] = v
		}
	} else if l == 2 {
		for k, _ := range ctx {
			delete(ctx, k)
		}
		for k, v := range res.Context {
			ctx[k] = v
		}
		for k, _ := range sts {
			delete(sts, k)
		}
		for k, v := range res.Status {
			sts[k] = v
		}
	}
}

//AddServant adds servant  for the service.
func (_obj *BackUp) AddServant(imp _impBackUp, obj string) {
	tars.AddServant(_obj, imp, obj)
}

//AddServant adds servant  for the service with context.
func (_obj *BackUp) AddServantWithContext(imp _impBackUpWithContext, obj string) {
	tars.AddServantWithContext(_obj, imp, obj)
}

type _impBackUp interface {
	Dump(DumpPath string, MirrorName string, Errmsg *string) (ret int32, err error)
	Restore(MirrorPath string, BinlogPath []string, LastTime int32, Normal bool, Errmsg *string) (ret int32, err error)
}
type _impBackUpWithContext interface {
	Dump(ctx context.Context, DumpPath string, MirrorName string, Errmsg *string) (ret int32, err error)
	Restore(ctx context.Context, MirrorPath string, BinlogPath []string, LastTime int32, Normal bool, Errmsg *string) (ret int32, err error)
}

func dump(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var DumpPath string
	err = _is.Read_string(&DumpPath, 1, true)
	if err != nil {
		return err
	}
	var MirrorName string
	err = _is.Read_string(&MirrorName, 2, true)
	if err != nil {
		return err
	}
	var Errmsg string
	if withContext == false {
		_imp := _val.(_impBackUp)
		ret, err := _imp.Dump(DumpPath, MirrorName, &Errmsg)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impBackUpWithContext)
		ret, err := _imp.Dump(ctx, DumpPath, MirrorName, &Errmsg)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = _os.Write_string(Errmsg, 3)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func restore(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var MirrorPath string
	err = _is.Read_string(&MirrorPath, 1, true)
	if err != nil {
		return err
	}
	var BinlogPath []string
	err, _, ty = _is.SkipToNoCheck(2, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		BinlogPath = make([]string, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = _is.Read_string(&BinlogPath[i0], 0, false)
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
	var LastTime int32
	err = _is.Read_int32(&LastTime, 3, true)
	if err != nil {
		return err
	}
	var Normal bool
	err = _is.Read_bool(&Normal, 4, true)
	if err != nil {
		return err
	}
	var Errmsg string
	if withContext == false {
		_imp := _val.(_impBackUp)
		ret, err := _imp.Restore(MirrorPath, BinlogPath, LastTime, Normal, &Errmsg)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impBackUpWithContext)
		ret, err := _imp.Restore(ctx, MirrorPath, BinlogPath, LastTime, Normal, &Errmsg)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = _os.Write_string(Errmsg, 5)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//Dispatch is used to call the server side implemnet for the method defined in the tars file. withContext shows using context or not.
func (_obj *BackUp) Dispatch(ctx context.Context, _val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
	_is := codec.NewReader(tools.Int8ToByte(req.SBuffer))
	_os := codec.NewBuffer()
	switch req.SFuncName {
	case "dump":
		err := dump(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "restore":
		err := restore(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(ctx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(ctx)
	if ok && c != nil {
		_context = c
	}
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(_os.ToBytes()),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}
	return nil
}
