package cmd

import (
	"DCacheClient/model/client"
	"DCacheClient/model/config"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cmdLineParse(ss []string) (k, v []string, e []int32) {
	for _, s := range ss {
		i := strings.Split(s, ":")
		k = append(k, i[0])
		v = append(v, i[1])
		t, _ := strconv.Atoi(i[2])
		e = append(e, int32(t))
	}
	return
}

func main() {
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

	cmd := os.Args[1]
	cmdList := os.Args[2:]

	switch cmd {
	case "get":
		key := cmdList[0]
		_, value, err := r.GetKV(key, opt)
		if err != nil {
			fmt.Sprintln(err)
		}
		fmt.Println(value)
	case "gets":
		keys := cmdList
		_, values, err := r.GetKVBatch(keys, opt)
		if err != nil {
			fmt.Sprintln(err)
		}
		fmt.Println(values)
	case "getall":
		index, err := strconv.Atoi(cmdList[0])
		if err != nil {
			fmt.Sprintln(err)
		}
		count, err := strconv.Atoi(cmdList[1])
		if err != nil {
			fmt.Sprintln(err)
		}
		_, value, _, err := r.GetAllKeys(int32(index), int32(count), opt)
		fmt.Println(value)
	case "check":
		keys := cmdList
		_, value, err := r.CheckKey(keys, opt)
		if err != nil {
			fmt.Sprintln(err)
		}
		fmt.Println(value)
	case "set":
		key := cmdList[0]
		value := cmdList[1]
		dirty := true
		ext, _ := strconv.Atoi(cmdList[2])
		ret, err := r.SetKV(key, value, dirty, int32(ext), opt)
		if err != nil {
			fmt.Sprintln(err)
		}
		fmt.Println(ret)
	case "insert":
		key := cmdList[0]
		value := cmdList[1]
		dirty := true
		ext, _ := strconv.Atoi(cmdList[2])
		ret, err := r.InsertKV(key, value, dirty, int32(ext), opt)
		if err != nil {
			fmt.Sprintln(err)
		}
		fmt.Println(ret)
	case "sets":
		// key-value example = key:value:0
		stringSlice := cmdList
		keys, values, exts := cmdLineParse(stringSlice)
		dirtys := func() []bool {
			d := make([]bool, len(keys))
			for i := 0; i < len(keys); i++ {
				d[i] = true
			}
			return d
		}()
		ret, m, err := r.SetKVBatch(keys, values, dirtys, exts, opt)
		if err != nil {
			fmt.Sprintln(err)
		}
		fmt.Println(ret, m)
	case "update":
		key := cmdList[0]
		value := cmdList[1]
		dirty := true
		op, _ := strconv.Atoi(cmdList[2])
		ext, _ := strconv.Atoi(cmdList[3])
		ret, s, err := r.UpdateKV(key, value, dirty, int32(op), int32(ext), opt)
		if err != nil {
			fmt.Sprintln(err)
		}
		fmt.Println(ret, s)
	case "del":
		key := cmdList[0]
		ret, err := r.DelKV(key, opt)
		if err != nil {
			fmt.Sprintln(err)
		}
		fmt.Println(ret)
	case "dels":
		keys := cmdList
		ret, m, err := r.DelKVBatch(keys, opt)
		if err != nil {
			fmt.Sprintln(err)
		}
		fmt.Println(ret, m)
	default:
		fmt.Println("Invalid options...")
	}
}