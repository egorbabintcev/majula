package core

import (
	"encoding/json"
	"majula/internal/infrastructure/storage"
)

type PackageStorage interface {
	GetPackage(name string) (storage.GetPackageRes, error)
	AddPackageVersion(name, version string, manifest json.RawMessage) (storage.AddPackageVersionRes, error)
	AddPackageTag(name, version, tag string) (storage.AddPackageTagRes, error)
}

type TarballStorage interface {
	SaveTarball(name, version string, tar []byte) (storage.SaveTarballRes, error)
	GetTarball(id string) (storage.GetTarballRes, error)
}

type Service struct {
	packageStorage PackageStorage
	tarballStorage TarballStorage
}

func NewService(ps PackageStorage, ts TarballStorage) *Service {
	return &Service{
		packageStorage: ps,
		tarballStorage: ts,
	}
}

func (s *Service) GetPackage(name string) (GetPackageRes, error) {
	p, err := s.packageStorage.GetPackage(name)

	if err != nil {
		return GetPackageRes{}, err
	}

	r := GetPackageRes{
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
	if _, err := s.packageStorage.AddPackageVersion(name, version, manifest); err != nil {
		return err
	}

	for _, t := range tags {
		if _, err := s.packageStorage.AddPackageTag(name, version, t); err != nil {
			return err
		}
	}

	if _, err := s.tarballStorage.SaveTarball(name, version, tar); err != nil {
		return err
	}

	return nil
}
