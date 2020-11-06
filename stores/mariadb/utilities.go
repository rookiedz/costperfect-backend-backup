package mariadb

import "time"

//CurrentDatetimeString ...
func CurrentDatetimeString() string {
	var datetime time.Time
	datetime = time.Now()
	return datetime.Format("2006-01-02 15:04:05")
}
