package main

import (
	"DCacheClient/model/client"
	"DCacheClient/model/config"
	"fmt"
	"os"
)

func main() {
	// 先展示一个简易的DEMO，其他功能后续添加
	basePath, err := os.Getwd()
	if err != nil {
		fmt.Sprintln(err)
	}
	cfg, err := config.ConfigParse(basePath)
	if err != nil {
		fmt.Sprintln(err)
	}
	r, _ := client.NewClient(cfg.Registry.Property, cfg.KVCache.CacheObj, cfg.KVCache.WCacheObj, cfg.KVCache.ModuleName)
	opt := make(map[string]string)

	// SET
	ret, err := r.SetKV("name", "likuo5",true, 0, opt)
	fmt.Println(ret)

	// GET
	value, err := r.GetKV("name", opt)
	fmt.Println(value)
}