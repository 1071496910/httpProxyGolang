package proxyServer

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/1071496910/httpProxyGolang/config"
)

type proxyServer struct {
}

type workProcess struct {
	Config config.Config
}

func NewWorkProcess(c config.Config) *workProcess {
	fmt.Printf("Config: %v\n", c)
	return &workProcess{
		Config: c,
	}
}

func (w *workProcess) GetBackend(serverName string, location string) string {

	for _, s := range w.Config.Servers {
		fmt.Printf("serverName:%v\n", serverName)
		for _, sn := range s.ServerName {
			if serverName == sn {
				fmt.Printf("%v==%v", serverName, sn)
				for _, l := range s.Locations {
					if l.Path == location {
						return "http://" +
							l.ProxyPass.Endporint[rand.Int()%len(l.ProxyPass.Endporint)]
					}
				}
			}
		}
	}
	return ""
}

func (w *workProcess) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	req, err := http.NewRequest(r.Method, w.GetBackend(r.Host, "/")+r.URL.Path, nil)
	if err != nil {
		fmt.Errorf("New Request err %v\n", err)
	}

	fmt.Printf("req:%v\n", req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("%v\n", err)
	}

	fmt.Printf("%v\n", resp)
	io.Copy(rw, resp.Body)

}

func (w *workProcess) Run() {
	http.ListenAndServe("127.0.0.1:80", w)
}
