package route

import (
	"fmt"
	"github.com/bytedance/sonic"
	Net "github.com/tobycroft/TuuzNet"
)

type route struct {
	Route string
}

func MainWsRouter() {
	for c := range Net.WsServer_ReadChannel {
		fmt.Println(c.Conn.RemoteAddr(), string(c.Message), c.Status)
		r := route{}
		err := sonic.Unmarshal(c.Message, &r)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch r.Route {
		case "login":
			Net.WsServer_WriteChannel <- c
			break
		}
	}
}
