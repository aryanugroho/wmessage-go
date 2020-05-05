package apiv0

import (
	"github.com/aryanugroho/advcorner/log"
	"github.com/aryanugroho/wmessage-go/pkg"
	"github.com/aryanugroho/wmessage-go/service"
	"github.com/aryanugroho/wmessage-go/util"
	"net/http"
)

func addMessage(w http.ResponseWriter, r *http.Request) {
	var payload pkg.Message
	err := util.BindJSON(r, &payload)
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = service.AddMessage(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
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

func listMessageSocket(w http.ResponseWriter, r *http.Request) {

}
