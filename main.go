package main

import (
	"fmt"
	"github.com/zserge/lorca"
	"os"
	"os/signal"
	"project/transmit/server"
	"sync"
)

//捕获异常
func recoverFromError() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}

func main() {
		var endWaiter sync.WaitGroup
		endWaiter.Add(1)
		end := make(chan interface{})
		go server.Run()
		go func(quit chan interface{}) {
			defer recoverFromError() //捕获异常
			ui, _ := lorca.New(fmt.Sprintf("http://127.0.0.1:27149/static/index.html"), "", 800, 600, "--disable-sync", " --disable-translate")
			defer ui.Close()
			quit<- (<-ui.Done()) //等待ui退出
		}(end)
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt)
		select {
		case <-signalChannel:
			endWaiter.Done()
		case <-end:
			endWaiter.Done()
		}
		endWaiter.Wait()

	//以下代码为不使用lorca包，直接调用电脑中chrom路径的方法
	//chBrowser := make(chan struct{})
	//chbackend := make(chan struct{})
	//chSignal := server.Listen()
	//go server.Run()
	//go server.GUI(chBrowser , chbackend)
	//// 当接收到中断信号时，执行下面的代码
	//for{
	//	select {
	//	case <-chSignal:
	//		chbackend <- struct{}{}
	//	case <-chBrowser:
	//		os.Exit(0)
	//	}
	//}
}
