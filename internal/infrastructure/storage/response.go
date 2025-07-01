package storage

import "encoding/json"

type GetPackumentRes struct {
	Name     string
	Versions map[string]json.RawMessage
	Tags     map[string]string
}

type AddPackumentVersionRes struct{}

type AddPackumentTagRes struct{}

type SaveTarballRes struct{
	Id string
}

type GetTarballRes struct {
	Content []byte
}
