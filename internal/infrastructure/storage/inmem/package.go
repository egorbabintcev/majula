package inmem

import (
	"encoding/json"
	"errors"
	"majula/internal/infrastructure/storage"
)

var (
	ErrPackageNotFound   = errors.New("failed to find package")
	ErrAddPackageVersion = errors.New("failed to save version manifest")
	ErrAddPackageTag     = errors.New("failed to save version tag")
)

type storageEntry struct {
	versions map[string]json.RawMessage
	tags     map[string]string
}

type PackageStorage struct {
	packages map[string]storageEntry
}

func NewPackageStorage() *PackageStorage {
	return &PackageStorage{
		packages: map[string]storageEntry{},
	}
}

func (s *PackageStorage) GetPackage(name string) (storage.GetPackageRes, error) {
	if _, exist := s.packages[name]; !exist {
		return storage.GetPackageRes{}, ErrPackageNotFound
	}

	p := s.packages[name]
	res := storage.GetPackageRes{
		Name:     name,
		Versions: p.versions,
		Tags:     p.tags,
	}

	return res, nil
}

func (s *PackageStorage) AddPackageVersion(name, version string, manifest json.RawMessage) (storage.AddPackageVersionRes, error) {
	if _, exist := s.packages[name]; !exist {
		s.packages[name] = storageEntry{
			versions: map[string]json.RawMessage{},
			tags:     map[string]string{},
		}
	}

	if _, exist := s.packages[name].versions[version]; exist {
		return storage.AddPackageVersionRes{}, ErrAddPackageVersion
	}

	s.packages[name].versions[version] = manifest

	return storage.AddPackageVersionRes{}, nil
}

func (s *PackageStorage) AddPackageTag(name, version, tag string) (storage.AddPackageTagRes, error) {
	if _, exist := s.packages[name]; !exist {
		s.packages[name] = storageEntry{
			versions: map[string]json.RawMessage{},
			tags:     map[string]string{},
		}
	}

	s.packages[name].tags[tag] = version

	return storage.AddPackageTagRes{}, nil
}
