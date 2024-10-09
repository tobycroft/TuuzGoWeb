package Net

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type WsData struct {
	Conn    *websocket.Conn
	Type    int
	Message []byte
	Status  bool
}

var WsServer_ReadChannel = make(chan WsData, 1)
var WsServer_WriteChannel = make(chan WsData, 1)

type WsServer struct {
	WsConfig *WsConfig
	url      string
	err      error
	Conn     *websocket.Conn
}

type WsConfig struct {
	PingReplyDelayInMs uint
	PongReplyDelayInMs uint
}

func (ws *WsServer) NewServer(w http.ResponseWriter, r *http.Request, responseHeader http.Header) {
	if ws.WsConfig == nil {
		ws.WsConfig = &WsConfig{
			PingReplyDelayInMs: 10,
		}
	}
	upd := websocket.Upgrader{}
	upd.EnableCompression = false
	upd.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws.Conn, ws.err = upd.Upgrade(w, r, responseHeader)
	if ws.err != nil {
		return
	}
	go ws.send_data()
	defer ws.Conn.Close()
	for {
		Type, message, err := ws.Conn.ReadMessage()
		switch Type {

		case websocket.TextMessage:
			WsServer_ReadChannel <- WsData{Conn: ws.Conn, Message: message, Type: Type, Status: true}
			break

		case websocket.BinaryMessage:
			WsServer_ReadChannel <- WsData{Conn: ws.Conn, Message: message, Type: Type, Status: true}
			break

		case websocket.PingMessage:
			go func() {
				time.Sleep(time.Duration(ws.WsConfig.PongReplyDelayInMs) * time.Millisecond)
				ws.Conn.WriteMessage(websocket.PongMessage, []byte("pong"))
			}()
			break

		case websocket.PongMessage:
			go func() {
				time.Sleep(time.Duration(ws.WsConfig.PingReplyDelayInMs) * time.Millisecond)
				ws.Conn.WriteMessage(websocket.PingMessage, []byte("ping"))
			}()
			break

		case websocket.CloseMessage, -1:
			go func() {
				select {
				case <-time.After(1 * time.Second):
					break
				case WsServer_WriteChannel <- WsData{Conn: ws.Conn, Message: message, Type: Type}:
					break
				}
			}()

			go func() {
				select {
				case <-time.After(1 * time.Second):
					break
				case WsServer_ReadChannel <- WsData{Conn: ws.Conn, Message: message, Type: Type}:
					break
				}
			}()
			return

		default:
			if err != nil {
				log.Println("server-read-error:", err)
				return
			}
			break
		}

	}
}

func (ws *WsServer) send_data() {
	for c := range WsServer_WriteChannel {
		switch c.Type {
		case websocket.TextMessage, websocket.BinaryMessage:
			err := c.Conn.WriteMessage(c.Type, c.Message)
			if err != nil {
				log.Println("server-send-error:", err)
				return
			}
			break

		case websocket.PingMessage:
			err := c.Conn.WriteMessage(websocket.PingMessage, []byte("ping"))
			if err != nil {
				log.Println("server-ping-error:", err)
				return
			}
			break

		case websocket.PongMessage:
			err := c.Conn.WriteMessage(websocket.PongMessage, []byte("pong"))
			if err != nil {
				log.Println("server-pong-error:", err)
				return
			}
			break

		case websocket.CloseMessage, -1:
			c.Conn.WriteMessage(websocket.CloseMessage, []byte("close"))
			return

		default:
			break
		}

	}
}
