package main

import (
	"fmt"
	//"gopkg.in/redis.v4"
	"github.com/go-redis/redis"
	"time"
)

func main() {
    var Timeout = 1 * time.Minute
    client := redis.NewClusterClient(&redis.ClusterOptions{
	    Addrs: 		[]string{
		    "10.80.137.192:6379",
			"10.80.137.193:6379",
			"10.80.137.194:6379",
			"10.80.137.192:6380",
			"10.80.137.193:6380",
			"10.80.137.194:6380"},
		Password: 	"nNnzvmX0MdoGIfyp",
	})
    err := client.Ping().Err()
	if err != nil {
	    fmt.Sprintln(err)
	}
	    fmt.Println(client.Set("foo", "bar", Timeout))
}	
