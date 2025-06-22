package web

import "encoding/json"

type ErrResBody struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

type GetWhoAmIRes struct {
	Username string `json:"username"`
}

type GetPkgRes struct {
	Name     string                     `json:"name"`
	Versions map[string]json.RawMessage `json:"versions"`
	Tags     map[string]string          `json:"dist-tags"`
}
