package filesystem

import (
	"errors"
	"log/slog"
	"majula/internal/infrastructure/storage"
	"os"
	"path/filepath"
)

var (
	ErrInitStorage = errors.New("failed to initialize storage")
	ErrSaveTarball = errors.New("failed to save tarball")
	ErrGetTarball  = errors.New("failed to get tarball")
)

type TarballStorage struct {
	logger   *slog.Logger
	location string
}

func NewTarballStorage(l *slog.Logger, base string) (*TarballStorage, error) {
	l = l.With(slog.String("component", "tarball_storage"))

	l.Info("Initializing tarball storage")

	info, err := os.Stat(base)

	if err != nil {
		return nil, ErrInitStorage
	}

	if !info.IsDir() {
		return nil, ErrInitStorage
	}

	return &TarballStorage{
		logger:   l,
		location: filepath.Join(base, "tarballs"),
	}, nil
}

func (s *TarballStorage) SaveTarball(name, version string, content []byte) (storage.SaveTarballRes, error) {
	id := name + "-" + version + ".tgz"

	path := filepath.Join(s.location, id)

	if _, err := os.Stat(path); err == nil || !errors.Is(err, os.ErrNotExist) {
		return storage.SaveTarballRes{}, ErrSaveTarball
	}

	err := os.WriteFile(path, content, 0644)

	if err != nil {
		return storage.SaveTarballRes{}, ErrSaveTarball
	}

	return storage.SaveTarballRes{
		Id: id,
	}, nil
}

func (s *TarballStorage) GetTarball(id string) (storage.GetTarballRes, error) {
	f, err := os.ReadFile(filepath.Join(s.location, id))

	if err != nil {
		return storage.GetTarballRes{}, ErrGetTarball
	}

	return storage.GetTarballRes{
		Content: f,
	}, nil
}
