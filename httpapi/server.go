package httpapi

type CoreServer struct {
	botIps []string
}

var Server *CoreServer

func NewCoreServer() *CoreServer {
	coreServer := &CoreServer{}
	Server = coreServer
	coreServer.botIps = make([]string, 100)
	return coreServer
}

func (c *CoreServer) Run() {
	e := InitEcho()
	e.Logger.Fatal(e.Start(":8888"))
}
