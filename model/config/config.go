package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Configure struct {
	Registry    registry
	KVCache     kvCache
}

type registry struct {
	Property    string
}

type kvCache struct {
	ModuleName	string
	WCacheObj	string
	CacheObj	string
}

var cfg Configure

func ConfigParse(basePath string) (Configure, error) {
	path := basePath + "/conf/conf.toml"
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return cfg, fmt.Errorf("decode file '%v' error, msg '%v'", path, err)
	}
	return cfg, nil
}