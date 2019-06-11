package client

import (
	DC "DCacheClient/pkg/DCache"
	"DCacheClient/pkg/myGoTools"
	"fmt"
)

func (r *Req) SetKV(key, value string, dirty bool, ext int32, opt map[string]string) (int32, error) {
	req := &DC.SetKVReq{
		ModuleName: r.ModuleName,
		Data: DC.SSetKeyValue{
			KeyItem: key,
			Value: value,
			Dirty: dirty,
			ExpireTimeSecond: ext,
		},
	}
	ret, err := WCacheApp.SetKV(req, opt)
	if err != nil {
		return ret, fmt.Errorf("ret '%v' and err '%v' as setting the key-value '%v, %v'",
			ret, err, key, value)
	}
	return ret, nil
}

func (r *Req) InsertKV(key, value string, dirty bool, ext int32, opt map[string]string) (int32, error) {
	req := &DC.SetKVReq{
		ModuleName: r.ModuleName,
		Data: DC.SSetKeyValue{
			KeyItem: key,
			Value: value,
			Dirty: dirty,
			ExpireTimeSecond: ext,
		},
	}
	ret, err := WCacheApp.InsertKV(req, opt)
	if err != nil {
		return ret, fmt.Errorf("ret '%v' and err '%v' as setting the key-value '%v, %v'",
			ret, err, key, value)
	}
	return ret, nil
}

func (r *Req) SetKVBatch(keys, values []string, dirtys []bool, exts []int32, opt map[string]string) (int32, map[string]int32, error) {
	var equal []int
	var sSetKV []DC.SSetKeyValue

	equal = append(equal, len(keys), len(values), len(dirtys), len(exts))
	ok, err := myGoTools.AllIntEqual(equal)
	if ! ok {
		return CheckElemErr, make(map[string]int32), fmt.Errorf("ssetkeyvalue element nums is not equal, err msg: %v",
			err)
	}
	for i := 0; i < equal[0]; i++ {
		sSetKV = append(sSetKV, DC.SSetKeyValue{
			KeyItem: 			keys[i],
			Value:				values[i],
			Dirty: 				dirtys[i],
			ExpireTimeSecond: 	exts[i],
		})
	}
	req := &DC.SetKVBatchReq{
		ModuleName: r.ModuleName,
		Data: sSetKV,
	}
	resp := new(DC.SetKVBatchRsp)
	ret, err := WCacheApp.SetKVBatch(req, resp, opt)
	if err != nil {
		return ret, resp.KeyResult, fmt.Errorf("ret '%v' and err '%v' as setting the KVS 'keys: %v', 'values: %v', RES '%v'",
			ret, err, keys, values, resp.KeyResult)
	}
	return ret, resp.KeyResult, nil
}

func (r *Req) UpdateKV(key, value string, dirty bool, op, ext int32, opt map[string]string) (int32, string, error){
	// options do not work, so should do check it
	// Answer: 只支持数字类型的value进行ADD/SUB/ADD_INSERT/SUB_INSERT操作 TODO 待解决
	req := &DC.UpdateKVReq{
		ModuleName: r.ModuleName,
		Data: DC.SSetKeyValue{
			KeyItem: key,
			Value: value,
			Dirty: dirty,
			ExpireTimeSecond: ext,
		},
		Option: DC.Op(op),
	}
	resp := new(DC.UpdateKVRsp)
	ret, err := WCacheApp.UpdateKV(req, resp, opt)
	if err != nil {
		return ret, resp.RetValue, fmt.Errorf("ret '%v' and err '%v' as updating the key-value '%v, %s'",
			ret, err, key, value)
	}
	return ret, resp.RetValue, nil
}

/*
func (r *WriteReq) EraseKV() {
	// Low utilization
}

func (r *WriteReq) EraseKVBatch() {
	// Low utilization
}
*/

func (r *Req) DelKV(key string, opt map[string]string) (int32, error) {
	req := &DC.RemoveKVReq{
		ModuleName: r.ModuleName,
		KeyInfo: DC.KeyInfo{
			KeyItem: key,
		},
	}
	ret, err := WCacheApp.DelKV(req, opt)
	if err != nil {
		return ret, fmt.Errorf("ret '%v' and err '%v' as deleting the key '%v'", ret, err, key)
	}
	return ret, nil
}

func (r *Req) DelKVBatch(keys []string, opt map[string]string) (int32, map[string]int32, error){
	var sKeyInfo []DC.KeyInfo
	for _, i := range keys {
		sKeyInfo = append(sKeyInfo, DC.KeyInfo{
			KeyItem: i,
		})
	}
	req := &DC.RemoveKVBatchReq{
		ModuleName: r.ModuleName,
		Data: sKeyInfo,
	}
	resp := new(DC.RemoveKVBatchRsp)
	ret, err := WCacheApp.DelKVBatch(req, resp, opt)
	if err != nil {
		return ret, resp.KeyResult, fmt.Errorf("ret '%v' and err '%v' as deleting the KVS 'keys: %v' RES '%v'", ret, err,
			keys, resp.KeyResult)
	}
	return ret, resp.KeyResult, nil
}