package web

import (
	"encoding/json"
	"net/http"
)

func respondWErr(w http.ResponseWriter, c int) {
	w.Header().Set("Content-Type", "application/json")

	res, _ := json.Marshal(ErrResBody{
		Code:    c,
		Message: http.StatusText(c),
	})

	w.WriteHeader(c)
	w.Write(res)
}
