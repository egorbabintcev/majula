package storage

import "encoding/json"

type GetPackageRes struct {
	Name     string
	Versions map[string]json.RawMessage
	Tags     map[string]string
}

type AddPackageVersionRes struct{}

type AddPackageTagRes struct{}

type SaveTarballRes struct{
	Id string
}

type GetTarballRes struct {
	Content []byte
}
