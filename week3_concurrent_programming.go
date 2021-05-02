package go_bootcamp

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func demo() {
	goConcurrentProgramming()
}

// goConcurrentProgramming 第三周并行编程作业
// 作业描述：基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func goConcurrentProgramming() {
	group, ctx := errgroup.WithContext(context.Background())
	group.Go(startHTTPServer)
	group.Go(processLinuxSignal)

	err := group.Wait()
	fmt.Println(err)
	fmt.Printf("error=%v", ctx.Err())
}

// startHTTPServer 启动 HTTP 服务器
func startHTTPServer() error {
	return errors.New("error occurs when starting http server")
}

// processLinuxSignal 处理 Linux 信号
func processLinuxSignal() error {
	// 接收 Linux 信号
	linuxKillSignal := make(chan os.Signal, 1)
	signal.Notify(linuxKillSignal, os.Interrupt)
	<-linuxKillSignal
	err := shutdownHTTPServer()
	if err != nil {
		return err
	}

	return errors.New("received kill signal from linux, exiting")
}

// shutdownHTTPServer 终止服务器进程
func shutdownHTTPServer() error {
	return errors.New("error occurs when shutting down http server")
}
