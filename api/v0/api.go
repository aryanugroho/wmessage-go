package apiv0

import (
	"fmt"
	"github.com/aryanugroho/wmessage-go/pkg"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

// Init is api registry
func Init() {
	// message related module route
	// POST
	http.HandleFunc("/v0/message/add", addMessage)

	// GET
	http.HandleFunc("/v0/message/list", listMessage)

	// WebSocket
	http.Handle("/v0/message/ws", websocket.Handler(func(conn *websocket.Conn) {
		func() {
			message := pkg.Message{}
			for {
				if err := websocket.JSON.Receive(conn, &message); err != nil {
					fmt.Println(err)
				} else {
					_, err := conn.Write([]byte(message.Text))
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}()
	}))

	// Run
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("failed to starting http server, err: ", err.Error())
	}
}
