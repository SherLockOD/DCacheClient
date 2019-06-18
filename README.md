# DCache Client Tools

## Config
Set config for the connecting DCache


### Registry
- property : set tars registry address   
`eg: tars.tarsregistry.QueryObj@tcp -h ip_addr -p 17890`


### KVCache
- ModuleName : set module name of the KVCache App   
`eg: DCacheApp_test`

 
- CacheObj : set Cache object name of the App   
`eg: DCache.DCacheApp_testKVCacheServer1-1.CacheObj`

    
- WCacheObj :  set WCache object name of the App    
`eg: DCache.DCacheApp_testKVCacheServer1-1.WCacheObj`

    
## Build
- build

    `bash build.sh build`

- clear

    `bash build.sh clear`     

## Quick Start
### Help
- `help`
```bash
./DCache-Cli help
```

### Read
- `get`
```bash
./DCache-Cli get key
```

- `gets`
```bash
./DCache-Cli gets key1 key2 key3 ...
```

- `getkeys`
```bash
./DCache-Cli getkeys
```

- `check`
```bash
./DCache-Cli check key1 key2 key3 ...
```

### Write
- `set`
```bash
./DCache-Cli set key value
```
- `sets`
```bash
./DCache-Cli sets key1:value1 key2:value2 key3:value3 ...
```

- `insert`
```bash
./DCache-Cli insert key value
```

- `del`
```bash
./DCache-Cli del key
```

- `dels`
```bash
./DCache-Cli dels key1 key2 key3
```

## Directory description
Directory | Description
---|---
conf | Configuration file of the DCache
model | Function models of the Client Tools
pkg | Proxy package of the DCache
test | Benchmark of the DCache
 
> demo.go is a demo script for the DCache Client
