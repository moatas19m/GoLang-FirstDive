package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseJSONBody(r *http.Request, v interface{}) error {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), v); err != nil {
			return err
		}
	}
	return nil
}
