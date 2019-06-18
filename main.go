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
		fmt.Sprintln(err)
	}
	cfg, err := config.ConfigParse(basePath)
	if err != nil {
		fmt.Sprintln(err)
	}
	r, _ := client.NewClient(cfg.Registry.Property, cfg.KVCache.CacheObj, cfg.KVCache.WCacheObj, cfg.KVCache.ModuleName)
	cmd.Execute(r)
}