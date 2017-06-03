package config

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
