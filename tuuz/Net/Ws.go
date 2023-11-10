package Net

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type WsClient struct {
	url          string
	Conn         *websocket.Conn
	retryTime    int
	retryLeft    int
	retryDelay   time.Duration
	retry        bool
	err          error
	ReadChannel  chan []byte
	WriteChannel chan []byte
}

func (ws WsClient) SetRetry(retry bool) WsClient {
	ws.retry = retry
	return ws
}

func (ws WsClient) SetRetryTime(retryTime int) WsClient {
	ws.retryTime = retryTime
	return ws
}

func (ws WsClient) SetRetryDelay(retryDelaySec time.Duration) WsClient {
	ws.retryDelay = retryDelaySec * time.Second
	return ws
}

func (ws WsClient) prepare_channel() {
	ws.ReadChannel = make(chan []byte, 1)
	ws.WriteChannel = make(chan []byte, 1)
}

func (ws WsClient) connect() (err error) {
	if ws.Conn != nil {
		ws.Conn.Close()
	} else {
		ws.prepare_channel()
	}
	ws.Conn, _, err = websocket.DefaultDialer.Dial(ws.url, nil)
	if err != nil {
		log.Println("Conn-Err:", err)
	}
	return
}

func (ws WsClient) NewConnect(url string) error {
	ws.url = url
	if ws.retry {
		if ws.retryDelay.Seconds() < 1 {
			ws.retryDelay = 5 * time.Second
		}
	}
	err := ws.connect()
	if err != nil {
		return err
	} else {
		go ws.recv_data()
		go ws.send_data()
	}
	return ws.err
}

func (ws WsClient) recv_data() {
	for {
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			ws.err = err
			log.Println("read:", err)
			return
		}
		ws.ReadChannel <- message
	}
}

func (ws WsClient) send_data() {
	for c := range ws.WriteChannel {
		err := ws.Conn.WriteMessage(websocket.TextMessage, c)
		if err != nil {
			ws.err = err
			log.Println("write:", err)
			return
		}
	}
}
