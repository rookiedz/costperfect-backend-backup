package api

//Entry ...
type Entry struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

//NewEntry ...
func NewEntry(status string, message string, data interface{}) Entry {
	return Entry{Status: status, Message: message, Data: data}
}

//NewEmptyEntry ...
func NewEmptyEntry(status string, message string) Entry {
	var data []string
	return Entry{Status: status, Message: message, Data: data}
}
