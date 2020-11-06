package response

import (
	"encoding/json"
	"net/http"
)

//Data ...
type Data struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

//NewData ...
func NewData(status string, message string, data interface{}) Data {
	return Data{Status: status, Message: message, Data: data}
}

//NewEmptyData ...
func NewEmptyData(status string, message string) Data {
	var empty []string
	return Data{Status: status, Message: message, Data: empty}
}

//JSON ...
func JSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
