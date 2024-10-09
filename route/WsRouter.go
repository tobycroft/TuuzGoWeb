package route

import (
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
)

func MainWsRouter() {
	for c := range Net.WsServer_ReadChannel {
		fmt.Println(c.Conn.RemoteAddr(), string(c.Message), c.Status)
		nd, err := sonic.Get(c.Message, "route")
		if err != nil {
			continue
		}
		r, err := nd.String()
		if err != nil {
			continue
		}
		switch r {
		case "login":
			break

		default:
			Net.WsServer_WriteChannel <- c
			break
		}
	}
}
