package Ws

import "github.com/gorilla/websocket"

var Clients map[*websocket.Conn]bool
var User_Client map[float64]*websocket.Conn
var Client_User map[*websocket.Conn]float64
