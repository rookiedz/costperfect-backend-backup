package mariadb

import (
	"fmt"
	"strings"
	"time"
)

//CurrentDatetimeString ...
func CurrentDatetimeString() string {
	var datetime time.Time
	datetime = time.Now()
	return datetime.Format("2006-01-02 15:04:05")
}

//ArrayInt64ToString ...
func ArrayInt64ToString(arr []int64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(arr), " ", delim, -1), "[]")
}
