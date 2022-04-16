

package main

import (
"fmt"
"time"
)


var LimitQueue map[string][]int64
var ok bool

// 时间滑动窗口限流
func LimitFreqSingle(queueName string, count uint, timeWindow int64) bool {
    currTime := time.Now().Unix()
    if LimitQueue == nil {
        LimitQueue = make(map[string][]int64)
    }
    if _, ok = LimitQueue[queueName]; !ok {
        LimitQueue[queueName] = make([]int64, 0)
    }
    //队列未满
    if uint(len(LimitQueue[queueName])) < count {
        LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
        return true
    }
    //队列满了,取出最早访问的时间
    earlyTime := LimitQueue[queueName][0]
    //说明最早期的时间还在时间窗口内,还没过期,所以不允许通过
    if currTime-earlyTime <= timeWindow {
        return false
    } else {
        //说明最早期的访问应该过期了,去掉最早期的
        LimitQueue[queueName] = LimitQueue[queueName][1:]
        LimitQueue[queueName] = append(LimitQueue[queueName], currTime)
    }
    return true
}

func limitIpFreq(timeWindow int64, count uint) bool {
    key := "limit:" +"127.0.0.1"
    if !LimitFreqSingle(key, count, timeWindow) {
        return false
    }
    return true
}

func main()  {

    for i := 0; i < 20; i++ {

        ok := limitIpFreq(10,5)
        if ok {
            fmt.Println("pass ", i)
        } else {
            fmt.Println("limit ", i)
        }
        time.Sleep(100 * time.Millisecond)
    }

}