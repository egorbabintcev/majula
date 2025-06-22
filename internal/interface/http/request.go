package web

import "encoding/json"

type PutPkgReqAttachments struct {
	ContentType string `json:"content_type"`
	Data        string `json:"data"`
	Length      int    `json:"length"`
}

type PutPkgReq struct {
	Name        string                          `json:"name"`
	Versions    map[string]json.RawMessage      `json:"versions"`
	Tags        map[string]string               `json:"dist-tags"`
	Attachments map[string]PutPkgReqAttachments `json:"_attachments"`
}
