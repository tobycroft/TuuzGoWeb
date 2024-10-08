package Net

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WsData struct {
	Conn    *websocket.Conn
	Message []byte
	Status  bool
}

var WsServer_ReadChannel = make(chan WsData, 1)
var WsServer_WriteChannel = make(chan WsData, 1)

type WsServer struct {
	url  string
	err  error
	Conn *websocket.Conn
}

func (ws *WsServer) NewServer(w http.ResponseWriter, r *http.Request, responseHeader http.Header) {
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
	for {
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			ws.err = err
			log.Println("server-read-error:", err)
			WsServer_ReadChannel <- WsData{Conn: ws.Conn, Message: message, Status: false}
			return
		}
		WsServer_ReadChannel <- WsData{Conn: ws.Conn, Message: message, Status: true}
	}
}

func (ws *WsServer) send_data() {
	for c := range WsServer_WriteChannel {
		err := c.Conn.WriteMessage(websocket.TextMessage, c.Message)
		if err != nil {
			log.Println("server-send-error:", err)
			return
		}
	}
}
