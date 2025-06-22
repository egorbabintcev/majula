package inmem

import (
	"encoding/json"
	"errors"
	"majula/internal/infrastructure/storage"
)

var (
	ErrPackumentNotFound   = errors.New("failed to find package")
	ErrAddPackumentVersion = errors.New("failed to save version manifest")
	ErrAddPackumentTag     = errors.New("failed to save version tag")
)

type storageEntry struct {
	versions map[string]json.RawMessage
	tags     map[string]string
}

type PackumentStorage struct {
	packages map[string]storageEntry
}

func NewPackumentStorage() *PackumentStorage {
	return &PackumentStorage{
		packages: map[string]storageEntry{},
	}
}

func (s *PackumentStorage) GetPackument(name string) (storage.GetPackumentRes, error) {
	if _, exist := s.packages[name]; !exist {
		return storage.GetPackumentRes{}, ErrPackumentNotFound
	}

	p := s.packages[name]
	res := storage.GetPackumentRes{
		Name:     name,
		Versions: p.versions,
		Tags:     p.tags,
	}

	return res, nil
}

func (s *PackumentStorage) AddPackumentVersion(name, version string, manifest json.RawMessage) (storage.AddPackumentVersionRes, error) {
	if _, exist := s.packages[name]; !exist {
		s.packages[name] = storageEntry{
			versions: map[string]json.RawMessage{},
			tags:     map[string]string{},
		}
	}

	if _, exist := s.packages[name].versions[version]; exist {
		return storage.AddPackumentVersionRes{}, ErrAddPackumentVersion
	}

	s.packages[name].versions[version] = manifest

	return storage.AddPackumentVersionRes{}, nil
}

func (s *PackumentStorage) AddPackumentTag(name, version, tag string) (storage.AddPackumentTagRes, error) {
	if _, exist := s.packages[name]; !exist {
		s.packages[name] = storageEntry{
			versions: map[string]json.RawMessage{},
			tags:     map[string]string{},
		}
	}

	s.packages[name].tags[tag] = version

	return storage.AddPackumentTagRes{}, nil
}
