package main

import (
	"github.com/1071496910/httpProxyGolang/config"
	"github.com/1071496910/httpProxyGolang/proxyServer"
)

func main() {
	c := config.GetConfig("config/config.json")
	w := proxyServer.NewWorkProcess(c)
	w.Run()
}
