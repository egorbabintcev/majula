package core

import (
	"encoding/json"
	"majula/internal/infrastructure/storage"
)

type PackumentStorage interface {
	GetPackument(name string) (storage.GetPackumentRes, error)
	AddPackumentVersion(name, version string, manifest json.RawMessage) (storage.AddPackumentVersionRes, error)
	AddPackumentTag(name, version, tag string) (storage.AddPackumentTagRes, error)
}

type TarballStorage interface {
	SaveTarball(id string, tar []byte) (storage.SaveTarballRes, error)
	GetTarball(id string) (storage.GetTarballRes, error)
}

type Service struct {
	packumentStorage PackumentStorage
	tarballStorage   TarballStorage
}

func NewService(ps PackumentStorage, ts TarballStorage) *Service {
	return &Service{
		packumentStorage: ps,
		tarballStorage:   ts,
	}
}

func (s *Service) GetPkg(name string) (GetPackumentRes, error) {
	p, err := s.packumentStorage.GetPackument(name)

	if err != nil {
		return GetPackumentRes{}, err
	}

	r := GetPackumentRes{
		Name:     p.Name,
		Versions: p.Versions,
		Tags:     p.Tags,
	}

	return r, nil
}

func (s *Service) GetTarball(id string) (GetTarballRes, error) {
	r, err := s.tarballStorage.GetTarball(id)

	if err != nil {
		return GetTarballRes{}, err
	}

	return GetTarballRes{
		Content: r.Content,
	}, nil
}

func (s *Service) PublishPkg(name, version string, tags []string, manifest json.RawMessage, tar []byte) error {
	if _, err := s.packumentStorage.AddPackumentVersion(name, version, manifest); err != nil {
		return err
	}

	for _, t := range tags {
		if _, err := s.packumentStorage.AddPackumentTag(name, version, t); err != nil {
			return err
		}
	}

	if _, err := s.tarballStorage.SaveTarball(name+"-"+version+".tgz", tar); err != nil {
		return err
	}

	return nil
}
