package main

import (
	"fmt"
	"log"
	"github.com/gorilla/websocket"
	"net/http"
	"os/exec"
	"io/ioutil"
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
	//log.Println("Succesfully upgraded connection")

	lastConn = conn
}

func RunCode(res http.ResponseWriter, req *http.Request) {
	source := req.FormValue("source")
	
	err := ioutil.WriteFile("source", []byte(source), 0755)
    if err != nil {
		fmt.Println("Error while writing file: ",err)
		lastConn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
        return;
    }
	
	cmd := exec.Command("sh", "-c", req.FormValue("cmd"))
	cmdout, err := cmd.CombinedOutput()
	
	if lastConn != nil {
		lastConn.WriteMessage(websocket.TextMessage, cmdout)
	}
}
