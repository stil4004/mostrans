package chat

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // оставим без изменений
var connections = []*websocket.Conn{}

func removeConn(slice []*websocket.Conn, val *websocket.Conn) []*websocket.Conn {
	index := -1
	for i, v := range slice {
		if v == val {
			index = i
			break
		}
	}

	if index != -1 {
		if index < len(slice)-1 {
			copy(slice[index:], slice[index+1:])
		}
		slice = slice[:len(slice)-1]
	}

	return slice
}

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Started WS")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	connections = append(connections, conn)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			connections = removeConn(connections, conn)
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Server: %s", message)
		for _, c := range connections {
			if c != conn {
				err = c.WriteMessage(messageType, message[:len(message)-1])
				if err != nil {
					connections = removeConn(connections, conn)
					log.Println("Error during message writing:", err)
					break
				}
			}
		}
	}
	connections = removeConn(connections, conn)
}

func Helloer(w http.ResponseWriter, r *http.Request) {
	body, err := r.GetBody()
	if err != nil {
		return
	}
	defer body.Close()
	var text []byte
	body.Read(text)
	fmt.Println(string(text))
	if r.Method == "GET" {
		fmt.Fprintln(w, "Пошел нахуй")
	}
}
