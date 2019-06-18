package main

import (
	"fmt"
	"sync"

	"DCacheClient/pkg/DCache"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/go-redis/redis"
	"testing"
)

var (
	// total request nums
	trn = 6000

	// goroutines nums
	gn = 1000

	// one goroutines request nums
	grn = trn / gn

	wg sync.WaitGroup
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
	b.N = gn
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
		go func() {
			for i := 0; i < grn; i++ {
				_, _ = app.SetKV(req, opt)
			}
		}()
	}
}

func Benchmark_DCacheGet(b *testing.B) {
	b.StopTimer()
	b.N = gn
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
		go func() {
			for i := 0; i < grn; i++ {
				_, err := app.GetKV(req, rsp, opt)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
		}()
	}
}

func Benchmark_RedisClusterSet(b *testing.B) {
	b.StopTimer()
	b.N = gn
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
		go func() {
			for i := 0; i < grn; i++ {
				client.Set("foo", "bar", 0)
			}
		}()
	}
}

func Benchmark_RedisClusterGet(b *testing.B) {
	b.StopTimer()
	b.N = gn
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
		go func() {
			for i := 0; i < grn; i++ {
				client.Get("foo")
			}
		}()
	}
}

func Benchmark_RedisSet(b *testing.B) {
	b.StopTimer()
	b.N = gn
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
		go func() {
			for i := 0; i < grn; i++ {
				client.Set("foo", "bar", 0)
			}
		}()
	}
}

func Benchmark_RedisGet(b *testing.B) {
	b.StopTimer()
	b.N = gn
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
		go func() {
			for i := 0; i < grn; i++ {
				client.Get("foo")
			}
		}()
	}
}
