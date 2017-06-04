package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Upstream struct {
	Timeout   int
	Endporint []string
}

type Location struct {
	Path      string
	ProxyPass Upstream
}

type Server struct {
	Port       []int
	ServerName []string
	Locations  []Location
}

type Config struct {
	MaxProcess int
	Servers    []Server
}

func GetConfig(configPath string) Config {

	fileBuff := make([]byte, 1024)
	//config := NewDefaultConfig()
	var config Config

	f, err := os.Open(configPath)
	if err != nil {
		fmt.Printf("Open file %s error: %v\n", configPath, err)
	}
	defer f.Close()
	count, err := f.Read(fileBuff)

	fmt.Println(f.Name())
	if err != nil {
		fmt.Printf("read file err %v\n", err)
	}

	fmt.Printf("read %d bytes: %q\n", count, fileBuff[:count])
	dec := json.NewDecoder(strings.NewReader(string(fileBuff)))
	err = dec.Decode(&config)
	//err = json.Unmarshal(fileBuff[:count], config)
	if err != nil {
		fmt.Printf("Parse config file err:%v\n", err)
	}
	return config
}

func EncodeDefaultConfig() {
	buff, err := json.Marshal(NewDefaultConfig())
	if err != nil {
		fmt.Printf("Decode json error :%v\n", err)
	}
	fmt.Printf(string(buff))

}

func NewDefaultConfig() Config {
	return Config{
		MaxProcess: 2,
		Servers: []Server{
			Server{
				Port: []int{
					80,
				},
				ServerName: []string{
					"www.baidu.com",
				},
				Locations: []Location{
					Location{
						Path: "/",
						ProxyPass: Upstream{
							Endporint: []string{
								//"127.0.0.1:8000",
								"14.215.177.37:80",
							},
							Timeout: 30,
						},
					},
				},
			},
		},
	}
}
