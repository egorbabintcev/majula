package inmem

import (
	"errors"
	"majula/internal/infrastructure/storage"
)

var (
	ErrSaveTarball     = errors.New("failed to save tarball")
	ErrTarballNotFound = errors.New("failed to find tarball")
)

type TarballStorage struct {
	tarballs map[string][]byte
}

func NewTarballStorage() *TarballStorage {
	return &TarballStorage{
		tarballs: map[string][]byte{},
	}
}

func (s *TarballStorage) SaveTarball(id string, tar []byte) (storage.SaveTarballRes, error) {
	if _, exist := s.tarballs[id]; exist {
		return storage.SaveTarballRes{}, ErrSaveTarball
	}

	s.tarballs[id] = tar

	return storage.SaveTarballRes{}, nil
}

func (s *TarballStorage) GetTarball(id string) (storage.GetTarballRes, error) {
	if _, exist := s.tarballs[id]; !exist {
		return storage.GetTarballRes{}, ErrTarballNotFound
	}

	t := s.tarballs[id]

	return storage.GetTarballRes{
		Content: t,
	}, nil
}
