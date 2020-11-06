package api

import (
	"errors"
	"strconv"
)

//ID64 ...
func ID64(i string) (int64, error) {
	if i == "" {
		return 0, errors.New("ID is empty")
	}
	return strconv.ParseInt(i, 10, 64)
}
