package main

import (
    "fmt"
    "github.com/go-redis/redis"
    "math/rand"
    "time"
)

//var (
//    redisdb *redis.Client
//)

func main() {
    rand.Seed(time.Now().Unix())
    redisdb := redis.NewClient(&redis.Options{
        Addr:     "192.168.159.133:6379",
        Password: "redis",
    })

    ping := redisdb.Ping()

    fmt.Printf("ping : %v\n", ping.String())

    //result, err := redisdb.Do("set", "test", "test-a").Result()
    //
    //fmt.Printf("result : %+v,err : %v\n",result,err)
    count := 500000

    for i := 0; i < count; i++ {
        strLen := i % 33
        if strLen < 20 {
            strLen = rand.Intn(30) + 10
        }
        //fmt.Printf("index : %d, str : %s,lenght : %v\n",i,RandString(20),strLen)

        redisKey := fmt.Sprintf("redis_key_%d", i)
        redisValue := RandString(strLen)
        redisdb.Do("set", redisKey, redisValue).Result()
    }

    fmt.Printf("set redis end\n")
}

var strByte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var strByteLen = len(strByte)

func RandString(length int) []byte {

    bytes := make([]byte, length)
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := 0; i < length; i++ {
        bytes[i] = strByte[r.Intn(strByteLen)]
    }

    return bytes
}
