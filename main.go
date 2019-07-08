package main

import (
	"DCacheClient/model/client"
	"DCacheClient/model/cmd"
	"DCacheClient/model/config"
	"DCacheClient/model/path"
	"fmt"
)

func main() {
	basePath, err := path.GetProjectPath(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	cfg, err := config.ConfigParse(basePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, _ := client.NewClient(cfg.Registry.Property, cfg.KVCache.CacheObj, cfg.KVCache.WCacheObj, cfg.KVCache.ModuleName)
	cmd.Execute(r)
}