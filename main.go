package main

import (
	"github.com/1071496910/httpProxyGolang/config"
	"github.com/1071496910/httpProxyGolang/proxyServer"
)

func main() {
	c := config.NewDefaultConfig()
	w := proxyServer.NewWorkProcess(c)
	w.Run()
}
