package apiv0

import (
	"fmt"
	"github.com/aryanugroho/wmessage-go/pkg"
	"github.com/aryanugroho/wmessage-go/service"
	"github.com/aryanugroho/wmessage-go/util"
	"golang.org/x/net/websocket"
	"net/http"
)

func addMessage(w http.ResponseWriter, r *http.Request) {
	var payload pkg.Message
	err := util.BindJSON(r, &payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = service.AddMessage(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn, err := websocket.Dial("ws://localhost:9000/v0/message/ws", "ws", "localhost:9000")
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := websocket.JSON.Send(conn, &payload); err != nil {
		// handle error
		fmt.Println(err.Error())
	}

	util.ToJSON(w, http.StatusCreated, `{"status": "OK"}`)
}

func listMessage(w http.ResponseWriter, r *http.Request) {
	list, err := service.GetMessages()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	util.ToJSON(w, http.StatusOK, list)
}
