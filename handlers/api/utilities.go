package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

//ID64 ...
func ID64(i string) (int64, error) {
	if i == "" {
		return 0, errors.New("ID is empty")
	}
	return strconv.ParseInt(i, 10, 64)
}

//JSON ...
func JSON(w http.ResponseWriter, status int, entry interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(entry)
}
