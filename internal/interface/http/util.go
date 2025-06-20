package web

import (
	"encoding/json"
	"net/http"
)

func respondWErr(w http.ResponseWriter, s int, m string) {
	w.Header().Set("Content-Type", "application/json")

	res, _ := json.Marshal(ErrResBody{
		Code:    s,
		Message: m,
	})

	w.WriteHeader(s)
	w.Write(res)
}
