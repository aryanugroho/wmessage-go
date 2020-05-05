package apiv0

import (
	"github.com/aryanugroho/wmessage-go/pkg"
	"github.com/aryanugroho/wmessage-go/service"
	"github.com/aryanugroho/wmessage-go/util"
	"net/http"
)

func addMessage(w http.ResponseWriter, r *http.Request) {
	var payload pkg.Message
	err := util.BindJSON(r, &payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = service.AddMessage(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	util.ToJSON(w, http.StatusCreated, `{"status": "OK"}`)
}

func listMessage(w http.ResponseWriter, r *http.Request) {

}

func listMessageSocket(w http.ResponseWriter, r *http.Request) {

}
