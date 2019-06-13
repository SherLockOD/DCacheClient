package main

import (
	//"DCacheClient/model/client"
	//"DCacheClient/model/config"
	"fmt"
	//"os"
	"DCacheClient/pkg/DCache"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/go-redis/redis"
	"testing"
)

/*
func Benchmark_Set(b *testing.B) {
	b.StopTimer()
	b.N = 1000
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
		go func() { _, _ = r.SetKV("foo", "bar", true, 0, opt) }()
	}
}

func Benchmark_Get(b *testing.B) {
	b.StopTimer()
	b.N = 100000
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
		func() { _, _ = r.GetKV("name", opt) }()
	}
}
*/
func Benchmark_DCacheSet(b *testing.B) {
	b.StopTimer()
	b.N = 11000
	comm := tars.NewCommunicator()
	obj := "DCache.DCacheApp_ABTestKVCacheServer1-1.WCacheObj"
	app := new(DCache.WCache)
	comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 10.80.137.193 -p 17890")
	comm.StringToProxy(obj, app)

	opt := make(map[string]string)

	sSetKeyValue := DCache.SSetKeyValue{
		KeyItem: "foo",
		Value:   "bar",
	}

	req := &DCache.SetKVReq{
		ModuleName: "DCacheApp_ABTest",
		Data:       sSetKeyValue,
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		go func() { _, _ = app.SetKV(req, opt) }()
	}
}

func Benchmark_DCacheGet(b *testing.B) {
	b.StopTimer()
	b.N = 10000
	comm := tars.NewCommunicator()
	obj := "DCache.DCacheApp_ABTestKVCacheServer1-1.CacheObj"
	app := new(DCache.Cache)
	comm.SetProperty("locator", "tars.tarsregistry.QueryObj@tcp -h 10.80.137.193 -p 17890")
	comm.StringToProxy(obj, app)

	opt := make(map[string]string)
	rsp := new(DCache.GetKVRsp)

	req := &DCache.GetKVReq{
		ModuleName: "DCacheApp_ABTest",
		KeyItem:    "foo",
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		go func() { _, _ = app.GetKV(req, rsp, opt) }()
	}
}

func Benchmark_RedisClusterSet(b *testing.B) {
	b.StopTimer()
	b.N = 1000000
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"10.80.137.192:6379",
			"10.80.137.193:6379",
			"10.80.137.194:6379",
			"10.80.137.192:6380",
			"10.80.137.193:6380",
			"10.80.137.194:6380"},
		Password: "nNnzvmX0MdoGIfyp",
	})
	err := client.Ping().Err()
	if err != nil {
		fmt.Sprintln(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		go client.Set("foo", "bar", 0)
	}
}

func Benchmark_RedisClusterGet(b *testing.B) {
	b.StopTimer()
	b.N = 1000000
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"10.80.137.192:6379",
			"10.80.137.193:6379",
			"10.80.137.194:6379",
			"10.80.137.192:6380",
			"10.80.137.193:6380",
			"10.80.137.194:6380"},
		Password: "nNnzvmX0MdoGIfyp",
	})
	err := client.Ping().Err()
	if err != nil {
		fmt.Sprintln(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		go client.Get("foo")
	}
}

func Benchmark_RedisSet(b *testing.B) {
	b.StopTimer()
	b.N = 100000
	client := redis.NewClient(&redis.Options{
		Addr:     "10.80.137.193:6381",
		Password: "",
		DB:       0,
	})
	err := client.Ping().Err()
	if err != nil {
		fmt.Sprintln(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		go client.Set("foo", "bar", 0)
	}
}

func Benchmark_RedisGet(b *testing.B) {
	b.StopTimer()
	b.N = 100000
	client := redis.NewClient(&redis.Options{
		Addr:     "10.80.137.193:6381",
		Password: "",
		DB:       0,
	})
	err := client.Ping().Err()
	if err != nil {
		fmt.Sprintln(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		b.ReportAllocs()
		go client.Get("foo")
	}
}
