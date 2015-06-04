package main

import (
	"fmt"
	"log"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
	"os/exec"
)


func main() {
	listenTo := "localhost:7777"
	
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/run", RunCode)
	http.HandleFunc("/ws", WebsocketHdl)

	fmt.Println("Serving http://"+listenTo)
	http.ListenAndServe(listenTo, nil)
}

var lastConn *websocket.Conn
func WebsocketHdl(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		log.Println(err)
		return
	} else if err != nil {
		log.Println(err)
		return
	}
	log.Println("Succesfully upgraded connection")

	lastConn = conn
}

func RunCode(res http.ResponseWriter, req *http.Request) {
	
	code := req.FormValue("code")

	f, _ := os.Create("source")
	f.WriteString(code)
	f.Sync()
	f.Close()
	
	cmd := exec.Command("sh", "-c", req.FormValue("cmd"))
	cmdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(res, err)
		return
	}
	
	if lastConn != nil {
		lastConn.WriteMessage(websocket.TextMessage, cmdout)
	}
}
