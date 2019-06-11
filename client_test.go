package main

import (
	"DCacheClient/model/client"
	"DCacheClient/model/config"
	"fmt"
	"os"
	"testing"
)

func Benchmark_Set(b *testing.B) {
	b.StopTimer()
	b.N = 10000
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
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		go func() {_, _ = r.SetKV("name", "ifeng",true, 10, opt)}()
	}
}

func Benchmark_Get(b *testing.B) {
	b.StopTimer()
	b.N = 10000
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
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		go func() {_, _ = r.GetKV("name", opt)}()
	}
}