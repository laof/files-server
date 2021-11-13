package controllers

import (
	"files-server/libs"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/julienschmidt/httprouter"
)

var talkList []Message = []Message{}
var userLogin = make(map[string]string)

type Message struct {
	Text   string    `json:"text"`
	Author string    `json:"author"`
	Type   string    `json:"type"`
	Time   time.Time `json:"time"`
}

type DataType struct {
	List    []Message `json:"list"`
	Success bool      `json:"success"`
}

func getCookie(s socketio.Conn) string {
	c := s.RemoteHeader().Get("Cookie")
	return c
}

func WebsocketServer(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	s.ServeHTTP(rw, r)
}

var s *socketio.Server

func init() {
	s = createServer()
	go s.Serve()
}

func createServer() *socketio.Server {
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

func GetTakHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	data := DataType{talkList, true}
	w.Write(libs.JsonData(data))
}
