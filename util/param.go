package util

import (
	"encoding/json"
	"net/http"
)

func BindJSON(req *http.Request, obj interface{}) error {
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(obj); err != nil {
		return err
	}

	return nil
}

func ToJSON(w http.ResponseWriter, statusCode int, obj interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(obj)
}
