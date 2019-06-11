package client

import (
	DC "DCacheClient/pkg/DCache"
	"fmt"
)

func (r *Req) GetKV(key string, opt map[string]string) (string, error) {
	req := &DC.GetKVReq{
		ModuleName: r.ModuleName,
		KeyItem: key,
	}
	resp := new(DC.GetKVRsp)
	ret, err := CacheApp.GetKV(req, resp, opt)
	if err != nil {
		return "", fmt.Errorf("ret '%v' and err '%v' as getting the key '%v'", ret, err, key)
	}
	return resp.Value, nil
}

func (r *Req) GetKVBatch(keys []string, opt map[string]string) ([]DC.SKeyValue, error) {
	req := &DC.GetKVBatchReq{
		ModuleName: r.ModuleName,
		Keys: keys,
	}
	resp := new(DC.GetKVBatchRsp)
	ret, err := CacheApp.GetKVBatch(req, resp, opt)
	if err != nil {
		return nil, fmt.Errorf("ret '%v' and err '%v' as getting the keys '%v'", ret, err, keys)
	}
	return resp.Values, nil
}

// 由 client 程序循环调用判断(resp.isEnd)
/*	keys := make([]string)
	keys = append(keys, resp.Keys...) //将后面的slice添加到前面slice里面
 */
func (r *Req) GetAllKeys(index, count int32, opt map[string]string) ([]string, bool, error) {
	req := &DC.GetAllKeysReq{
		ModuleName: r.ModuleName,
		Index: index,
		Count: count,
	}
	resp := new(DC.GetAllKeysRsp)
	ret, err := CacheApp.GetAllKeys(req, resp, opt)
	if err != nil {
		return nil, false, fmt.Errorf("ret '%v' and err '%v' as getting the index of range '%v - %v'", ret, err, index, count)
	}
	return resp.Keys, resp.IsEnd, nil
}


func (r *Req) CheckKey(keys []string, opt map[string]string) (string, error) {
	req := &DC.CheckKeyReq{
		ModuleName: r.ModuleName,
		Keys: keys,
	}
	resp := new(DC.CheckKeyRsp)
	ret, err := CacheApp.CheckKey(req, resp, opt)
	if err != nil {
		return "", fmt.Errorf("ret '%v' and err '%v' as getting the keys '%v'", ret, err, keys)
	}
	return fmt.Sprintf("%+v", resp.KeyStat), nil
}