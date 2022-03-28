package main

import (
    "context"
    "fmt"
    "golang.org/x/sync/errgroup"
    "io"
    "net/http"
    "os"
    "os/signal"
)

func startHttpServer(svr *http.Server) error {
    fmt.Println("server start")
    http.HandleFunc("/hello",helloServer)
    return svr.ListenAndServe()

}

func helloServer(w http.ResponseWriter,req *http.Request)  {
    io.WriteString(w,"hello geek。。。")
}

func main() {
    ctx := context.Background()

    ctx, cancel := context.WithCancel(ctx)

    group, errCtx := errgroup.WithContext(ctx)

    svr := &http.Server{Addr: ":9898"}


    group.Go(func() error {
        // 启动服务
        return startHttpServer(svr)
    })

    group.Go(func() error { // 停止服务
        <-errCtx.Done()
        fmt.Println("server stop")
        return svr.Shutdown(errCtx)
    })

    chanel := make(chan os.Signal, 1)
    signal.Notify(chanel) // 监听信号量

    group.Go(func() error {
        for  {
            select {
            case <-errCtx.Done():
                fmt.Println("error : ",errCtx.Err())
                return errCtx.Err() // 监听到上下文异常
            case <-chanel:
                fmt.Println("cancel")
                cancel() // 接收到信号量

            }
        }
    })

    if err := group.Wait(); err !=nil{
        fmt.Println("group error : ",err)
    }

    fmt.Println("server exit")

}
