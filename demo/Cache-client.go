package main

import (
	"DCacheClient/pkg/DCache"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
)

var comm *tars.Communicator

func main() {
	comm = tars.NewCommunicator()
	obj := "DCache.DCacheApp_ABTestKVCacheServer1-1.CacheObj"
	app := new(DCache.Cache)
	comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 10.80.137.193 -p 17890")
	comm.StringToProxy(obj, app)

	rsp := new(DCache.GetKVRsp)

	opt := make(map[string]string)
	req := &DCache.GetKVReq{
		ModuleName: "DCacheApp_ABTest",
		KeyItem: 	"name",
	}

	ret, err := app.GetKV(req, rsp, opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ret: ", ret, "value: ", rsp.Value)
}
