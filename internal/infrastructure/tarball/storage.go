package tarball

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"
)

type Storage struct {
	logger   *slog.Logger
	location string
}

func NewStorage(l *slog.Logger, base string) (*Storage, error) {
	l = l.With(slog.String("component", "tarball_storage"))

	l.Info("Initializing tarball storage")

	info, err := os.Stat(base)

	if err != nil {
		return nil, ErrNoLocationDir
	}

	if !info.IsDir() {
		return nil, ErrNoLocationDir
	}

	return &Storage{
		logger:   l,
		location: filepath.Join(base, "tarballs"),
	}, nil
}

func (s *Storage) Save(name, version string, content []byte) (SaveResponse, error) {
	id := name + "-" + version + ".tgz"

	path := filepath.Join(s.location, id)

	_, err := os.Stat(path)

	if _, err := os.Stat(path); err == nil {
		return SaveResponse{}, ErrAlreadyExist
	} else if !errors.Is(err, os.ErrNotExist) {
		return SaveResponse{}, ErrNotFound
	}

	err = os.WriteFile(path, content, 0644)

	if err != nil {
		return SaveResponse{}, ErrWrite
	}

	return SaveResponse{
		Id: id,
	}, nil
}

func (s *Storage) Get(id string) (GetResponse, error) {
	f, err := os.ReadFile(filepath.Join(s.location, id))

	if err != nil {
		return GetResponse{}, ErrNotFound
	}

	return GetResponse{
		Tarball: f,
	}, nil
}
