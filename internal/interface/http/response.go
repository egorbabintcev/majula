package web

type ErrResBody struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
