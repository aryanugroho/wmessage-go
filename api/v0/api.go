package apiv0

import "net/http"

func Init() {
	// message related module route
	// POST
	http.HandleFunc("/v0/message/add", addMessage)

	// GET
	http.HandleFunc("/v0/message/list", listMessage)

	// WebSocket
	http.HandleFunc("/v0/message/ws", listMessageSocket)
}
