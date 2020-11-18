package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
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

//FieldsExist ...
func FieldsExist(s interface{}) {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		vf := v.Field(i)
		switch sf.Type.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
			// if vf.IsNil() {
			// 	log.Println(sf.Name)
			// }
			break
		case reflect.Struct:
			//log.Printf(sf.Name)
			FieldsExist(vf.Interface())
			break
		default:
			if vf.Interface() != reflect.Zero(sf.Type).Interface() {
				if vf == reflect.Zero(sf.Type) {
					//log.Println(sf.Name)
				}
				log.Println(sf.Name)
			}
			//log.Println(sf.Name)
			//log.Println(vf)

			//log.Println(reflect.Zero(sf.Type))
			//log.Println(vf.Interface())
		}
	}
}
