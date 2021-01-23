package api

//Data ...
type Data struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
	Total   int64       `json:"total,omitempty"`
}

//EmptyData ...
type EmptyData struct{}

//NewEmptyData ...
func NewEmptyData() EmptyData {
	return EmptyData{}
}

//Success ...
func Success(data interface{}) Data {
	return Data{Status: "success", Data: data}
}

//Total ...
func Total(total int64, data interface{}) Data {
	return Data{Status: "success", Total: total, Data: data}
}

//NotFound ...
func NotFound() Data {
	return Data{Status: "success", Data: EmptyData{}}
}

//Failure ...
func Failure(err error) Data {
	return Data{Status: "failure", Message: err.Error(), Data: EmptyData{}}
}

//Err ...
func Err(err error) Data {
	return Data{Status: "error", Message: err.Error(), Data: EmptyData{}}
}
