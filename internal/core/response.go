package core

import "encoding/json"

type GetPackageRes struct {
	Name     string
	Versions map[string]json.RawMessage
	Tags     map[string]string
}

type GetTarballResponse struct {
	Tarball []byte
}
