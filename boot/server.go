package boot

import (
	"mermaid/global"
)

func bootHTTP() error {
	serverHost := global.Conf.Server.Addr
	return global.InitHttp(serverHost)
}
