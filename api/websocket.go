package api

import (
	"encoding/json"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

var talkList []Message = []Message{}
var userTotal = 0
var userLogin = make(map[string]string)

type Message struct {
	Text   string    `json:"text"`
	Author string    `json:"author"`
	Type   string    `json:"type"`
	Time   time.Time `json:"time"`
}

func getCookie(s socketio.Conn) string {
	c := s.RemoteHeader().Get("Cookie")
	return c
}

func WebsocketServer() *socketio.Server {

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		return nil
	})

	server.OnEvent("/", "connection succeeded", func(s socketio.Conn, wsid string) {
		ck := getCookie(s)
		if userLogin[ck] == "" {
			userLogin[ck] = wsid
		}
	})

	server.OnEvent("/", "chat message", func(s socketio.Conn, txt string) {

		ck := getCookie(s)

		mes := Message{
			txt,
			userLogin[ck],
			"message",
			time.Now()}

		talkList = append(talkList, mes)
		server.BroadcastToNamespace("/", "chat message", mes)

	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	return server
}

type DataType struct {
	List    []Message `json:"list"`
	Success bool      `json:"success"`
}

func GetTakHistory() []byte {

	data := DataType{talkList, true}
	v, _ := json.Marshal(data)

	return v
}
