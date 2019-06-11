//Package DCache comment
// This file war generated by tars2go 1.1
// Generated from BinLog.tars
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

//BinLog struct
type BinLog struct {
	s m.Servant
}

//GetLog is the proxy function for the method defined in the tars file, with the context
func (_obj *BinLog) GetLog(Req *BinLogReq, Rsp *BinLogRsp, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = Req.WriteBlock(_os, 1)
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
	err = _obj.s.Tars_invoke(ctx, 0, "getLog", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*Rsp).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//GetLogWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *BinLog) GetLogWithContext(ctx context.Context, Req *BinLogReq, Rsp *BinLogRsp, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = Req.WriteBlock(_os, 1)
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
	err = _obj.s.Tars_invoke(ctx, 0, "getLog", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*Rsp).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//GetLogCompress is the proxy function for the method defined in the tars file, with the context
func (_obj *BinLog) GetLogCompress(Req *BinLogReq, Rsp *BinLogRsp, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = Req.WriteBlock(_os, 1)
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
	err = _obj.s.Tars_invoke(ctx, 0, "getLogCompress", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*Rsp).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//GetLogCompressWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *BinLog) GetLogCompressWithContext(ctx context.Context, Req *BinLogReq, Rsp *BinLogRsp, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = Req.WriteBlock(_os, 1)
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
	err = _obj.s.Tars_invoke(ctx, 0, "getLogCompress", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*Rsp).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//GetLastBinLogTime is the proxy function for the method defined in the tars file, with the context
func (_obj *BinLog) GetLastBinLogTime(LastTime *uint32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()

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
	err = _obj.s.Tars_invoke(ctx, 0, "getLastBinLogTime", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_uint32(&(*LastTime), 1, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//GetLastBinLogTimeWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *BinLog) GetLastBinLogTimeWithContext(ctx context.Context, LastTime *uint32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "getLastBinLogTime", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_uint32(&(*LastTime), 1, true)
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
func (_obj *BinLog) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *BinLog) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}
func (_obj *BinLog) setMap(l int, res *requestf.ResponsePacket, ctx map[string]string, sts map[string]string) {
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
func (_obj *BinLog) AddServant(imp _impBinLog, obj string) {
	tars.AddServant(_obj, imp, obj)
}

//AddServant adds servant  for the service with context.
func (_obj *BinLog) AddServantWithContext(imp _impBinLogWithContext, obj string) {
	tars.AddServantWithContext(_obj, imp, obj)
}

type _impBinLog interface {
	GetLog(Req *BinLogReq, Rsp *BinLogRsp) (ret int32, err error)
	GetLogCompress(Req *BinLogReq, Rsp *BinLogRsp) (ret int32, err error)
	GetLastBinLogTime(LastTime *uint32) (ret int32, err error)
}
type _impBinLogWithContext interface {
	GetLog(ctx context.Context, Req *BinLogReq, Rsp *BinLogRsp) (ret int32, err error)
	GetLogCompress(ctx context.Context, Req *BinLogReq, Rsp *BinLogRsp) (ret int32, err error)
	GetLastBinLogTime(ctx context.Context, LastTime *uint32) (ret int32, err error)
}

func getLog(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var Req BinLogReq
	err = Req.ReadBlock(_is, 1, true)
	if err != nil {
		return err
	}
	var Rsp BinLogRsp
	if withContext == false {
		_imp := _val.(_impBinLog)
		ret, err := _imp.GetLog(&Req, &Rsp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impBinLogWithContext)
		ret, err := _imp.GetLog(ctx, &Req, &Rsp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = Rsp.WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func getLogCompress(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var Req BinLogReq
	err = Req.ReadBlock(_is, 1, true)
	if err != nil {
		return err
	}
	var Rsp BinLogRsp
	if withContext == false {
		_imp := _val.(_impBinLog)
		ret, err := _imp.GetLogCompress(&Req, &Rsp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impBinLogWithContext)
		ret, err := _imp.GetLogCompress(ctx, &Req, &Rsp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = Rsp.WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func getLastBinLogTime(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var LastTime uint32
	if withContext == false {
		_imp := _val.(_impBinLog)
		ret, err := _imp.GetLastBinLogTime(&LastTime)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impBinLogWithContext)
		ret, err := _imp.GetLastBinLogTime(ctx, &LastTime)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = _os.Write_uint32(LastTime, 1)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//Dispatch is used to call the server side implemnet for the method defined in the tars file. withContext shows using context or not.
func (_obj *BinLog) Dispatch(ctx context.Context, _val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
	_is := codec.NewReader(tools.Int8ToByte(req.SBuffer))
	_os := codec.NewBuffer()
	switch req.SFuncName {
	case "getLog":
		err := getLog(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "getLogCompress":
		err := getLogCompress(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "getLastBinLogTime":
		err := getLastBinLogTime(ctx, _val, _os, _is, withContext)
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
