package main

import (
	"DCacheClient/pkg/DCache"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
)

var comm *tars.Communicator

func setDemo() {
	comm = tars.NewCommunicator()
	obj := "DCache.DCacheApp_ABTestKVCacheServer1-1.WCacheObj"
	app := new(DCache.WCache)
	comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 10.80.137.193 -p 17890")
	comm.StringToProxy(obj, app)

	opt := make(map[string]string)

	sSetKeyValue := DCache.SSetKeyValue{
		KeyItem:	"namex",
		Value:		"ifeng",
	}

	req := &DCache.SetKVReq{
		ModuleName: "DCacheApp_ABTest",
		Data: 		sSetKeyValue,
	}

	fmt.Println(*req)
	ret, err := app.SetKV(req, opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ret: ", ret)
}

func updateDemo() {
	comm = tars.NewCommunicator()
	obj := "DCache.DCacheApp_ABTestKVCacheServer1-1.WCacheObj"
	app := new(DCache.WCache)
	comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 10.80.137.193 -p 17890")
	comm.StringToProxy(obj, app)

	opt := make(map[string]string)

	sSetKeyValue := DCache.SSetKeyValue{
		KeyItem:	"name",
		Value:		"ifeng",
		Dirty:		true,
		ExpireTimeSecond: 0,
	}

	req := &DCache.UpdateKVReq{
		ModuleName: "DCacheApp_ABTest",
		Data: 		sSetKeyValue,
		Option:     0,
	}

	resp := new(DCache.UpdateKVRsp)
	ret, err := app.UpdateKV(req, resp, opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ret: ", ret, resp.RetValue, err)
}

func main() {
	// setDemo()
	updateDemo()
}