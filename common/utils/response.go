package utils

import (
	"encoding/json"
	"net/http"
)

// OutputJSON func used for return response in JSON if success or return error when fail
func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
