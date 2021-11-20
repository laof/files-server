package controllers

import (
	"encoding/json"
	"files-server/utils"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

const (
	Sign = "sign"
	Chat = "chat"
)

type Client struct {
	send chan []byte
	mt   int
	conn *websocket.Conn
	id   string
}

func (client *Client) sendMessage(m Message) error {
	mt := client.mt
	e := BroadcastMessage{Chat, m}
	return client.conn.WriteMessage(mt, utils.JsonData(e))
}

type BroadcastMessage struct {
	Type string  `json:"type"`
	Data Message `json:"data"`
}
type RegisterEvent struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

var talkList []Message = []Message{}

type Message struct {
	Text   string    `json:"text"`
	Author string    `json:"author"`
	Type   string    `json:"type"`
	Time   time.Time `json:"time"`
}

type Hub struct {
	clients map[string]*Client

	broadcast chan Message

	register chan *Client

	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.id] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.id]; ok {
				delete(h.clients, client.id)
				close(client.send)
			}
		case m := <-h.broadcast:
			for _, client := range h.clients {
				client.sendMessage(m)
			}
		}
	}
}

type Socket struct{}

var socket = Socket{}

var hub = newHub()

func init() {
	go hub.run()
}

func (s *Socket) read(txt []byte) *RegisterEvent {
	event := new(RegisterEvent)

	err := json.Unmarshal(txt, &event)

	if err != nil {
		return new(RegisterEvent)
	}

	return event
}

var upgrader = websocket.Upgrader{}

func ChatServer(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	handleFunc(rw, r)
}

func handleFunc(w http.ResponseWriter, r *http.Request) {

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	var send *Client

	defer func() {

		if send.id != "" {
			hub.unregister <- send
		}

		c.Close()
	}()

	var userid = ""

	for {

		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}

		e := socket.read(message)

		if e.Type == Sign {
			userid = e.Data
			send = &Client{make(chan []byte), mt, c, e.Data}
			hub.register <- send
			continue
		}

		mes := Message{
			e.Data,
			userid,
			"message",
			time.Now()}

		talkList = append(talkList, mes)

		hub.broadcast <- mes

	}
}

type historyDataType struct {
	List    []Message `json:"list"`
	Success bool      `json:"success"`
}

func GetTakHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	data := historyDataType{talkList, true}
	w.Write(utils.JsonData(data))
}
