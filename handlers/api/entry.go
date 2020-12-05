package api

//Data ...
type Data struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

//EmptyData ...
type EmptyData struct{}

//NewEmptyData ...
func NewEmptyData() EmptyData {
	return EmptyData{}
}

//Success ...
func Success(endpoint string, data interface{}) Data {
	var entry map[string]interface{}
	entry = make(map[string]interface{})
	entry[endpoint] = data
	return Data{Status: "success", Data: entry}
}

//Done ...
func Done(endpoint string, data interface{}) Data {
	var entry map[string]interface{}
	entry = make(map[string]interface{})
	entry[endpoint] = data
	return Data{Status: "done", Data: entry}
}

//Total ...
func Total(endpoint string, total int64, data interface{}) Data {
	var entry map[string]interface{}
	entry = make(map[string]interface{})
	entry[endpoint] = data
	entry["total"] = total
	return Data{Status: "success", Data: entry}
}

//NotFound ...
func NotFound(endpoint string) Data {
	var data map[string]EmptyData
	data = make(map[string]EmptyData)
	data[endpoint] = EmptyData{}
	return Data{Status: "success", Data: data}
}

//Failure ...
func Failure(endpoint string, err error) Data {
	var data map[string]EmptyData
	data = make(map[string]EmptyData)
	data[endpoint] = EmptyData{}
	return Data{Status: "failure", Message: err.Error(), Data: data}
}

//Err ...
func Err(endpoint string, err error) Data {
	var data map[string]EmptyData
	data = make(map[string]EmptyData)
	data[endpoint] = EmptyData{}
	return Data{Status: "error", Message: err.Error(), Data: data}
}
