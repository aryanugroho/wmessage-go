package apiv0

import (
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
	http.HandleFunc("/v0/message/ws", listMessageSocket)

	// Run
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("failed to starting http server, err: ", err.Error())
	}
}
