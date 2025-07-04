package tarball

import "errors"

var (
	ErrNoLocationDir = errors.New("failed to find location directory")
	ErrNotFound = errors.New("failed to find tarball")
	ErrAlreadyExist = errors.New("tarball already exist")
	ErrWrite = errors.New("failed to write tarball file")
)