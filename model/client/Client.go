package client

import (
	DC "DCacheClient/pkg/DCache"
	"github.com/TarsCloud/TarsGo/tars"
)

type Req struct {
	ModuleName	string
}

var (
	comm 		*tars.Communicator
	WCacheApp	*DC.WCache
	CacheApp	*DC.Cache
)

// TODO 后期如果有连接其他obj的需求，可以考虑把obj放到slice里，循环导入
func NewClient(prop, cObj, wObj, module string) (*Req, error) {
	comm = tars.NewCommunicator()
	comm.SetProperty("locator", prop)
	WCacheApp = new(DC.WCache)
	CacheApp = new(DC.Cache)
	comm.StringToProxy(wObj, WCacheApp)
	comm.StringToProxy(cObj, CacheApp)
	return &Req{
		ModuleName: module,
	}, nil
}