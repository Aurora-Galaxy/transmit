package server

import (
	"os"
	"os/exec"
	"os/signal"
)

const port string = "27149"

func GUI(chBrowserDie chan struct{} , chBackendDie chan struct{}) {
	chormePath := "C:\\Users\\hp-pc\\AppData\\Local\\Google\\Chrome\\Application\\chrome.exe"
	cmd := exec.Command(chormePath,"--app=http://127.0.0.1:" + port +"/static/index.html")
	cmd.Start()
	go func() {  //后台关闭
		<-chBackendDie
		cmd.Process.Kill()
	}()
	go func(){
		cmd.Wait() //等待用户关闭浏览器
		chBrowserDie <- struct{}{}
	}()

}

func Listen()chan os.Signal{
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)
	return chSignal
}