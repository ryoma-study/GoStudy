package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

/**
题目：基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

思考：基于题干信息，拆分三个任务：
1. 实现 HTTP server 的启动和关闭
2. 基于 errgroup 实现多个 goroutine 的级联退出：保证全部注销成功
3. 监听 linux signal 信号：比如支持 kill -9 或 Ctrl+C
*/

func defaultHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello in %s!", server.Addr)
}

func showdownHandler(_ http.ResponseWriter, _ *http.Request) {
	shutdownChan <- struct{}{}
}

func startHttpServer(server *http.Server) error {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/shutdown", showdownHandler)

	fmt.Printf("http server start in %s\n", server.Addr)
	err := server.ListenAndServe()
	return err
}

var server *http.Server
var shutdownChan = make(chan struct{})

func main() {
	// http server
	server = &http.Server{Addr: ":8080"}

	fmt.Printf("Before: %d\n", runtime.NumGoroutine())
	// cancel 准备级联取消
	ctx, cancel := context.WithCancel(context.Background())

	// errgroup 的取消
	group, errCtx := errgroup.WithContext(ctx)

	group.Go(func() error {
		fmt.Println("g1 start")
		err := startHttpServer(server)
		if err != nil {
			fmt.Println("g1 will exit.", err.Error())
		}

		return err
	})

	/**
	调用 /shutdown 接口时, shutdownChan 写入数据:
	(1) g2 执行 server.shutdown, 此时 g1 httpServer 会退出
	(2) g1 退出后，会将 error 加到 errgroup
	(3) g3 中会执行 case errCtx.Done, context 将不再阻塞，g3 会随之退出
	(4) main 函数中的 g.Wait() 退出，所有协程都会退出
	*/
	/**
	g2 receive /shutdown
	g2 http server will be stopped
	g1 will exit. http: Server closed
	g3 get err: context canceled
	errgroup err:  http: Server closed
	*/
	group.Go(func() error {
		fmt.Println("g2 start")
		select {
		case <-errCtx.Done():
			fmt.Printf("g2 get err: %v\n", ctx.Err().Error())
		case <-shutdownChan:
			fmt.Println("g2 receive /shutdown")
		}

		fmt.Println("g2 http server will be stopped")
		return server.Shutdown(ctx)
	})

	/**
	比如执行 -9 信号时，terminalChan 收到数据:
	(1) 执行 cancel()，g3 退出
	(2) g2 中会执行 case errCtx.Done，随之执行 Shutdown 后 g2 退出
	(3) g1 被 Shutdown 直接触发 http: Server closed 进而导致关闭
	*/
	/**
	g3 get os signal: terminated
	g2 get err: context canceled
	g2 http server will be stopped
	g1 will exit. http: Server closed
	errgroup err:  context canceled
	*/
	group.Go(func() error {
		fmt.Println("g3 start")

		terminalChan := make(chan os.Signal, 1)
		signal.Notify(terminalChan, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-errCtx.Done():
			fmt.Printf("g3 get err: %v\n", errCtx.Err().Error())
		case sig := <-terminalChan:
			fmt.Printf("g3 get os signal: %v\n", sig)
			defer cancel()
		}

		return nil
	})

	fmt.Printf("now: %d\n", runtime.NumGoroutine())
	if err := group.Wait(); err != nil {
		fmt.Println("errgroup err: ", err)
	}

	fmt.Printf("After: %d\n", runtime.NumGoroutine())
}
