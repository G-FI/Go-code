package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func init() {
	log.SetPrefix("\x1b[92m[GO-BOOKSTORE-SERVER]\x1b[0m")
}

func Logf(format string, v ...any) {

	log.Printf(format, v...)
}

func ParseBody(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
