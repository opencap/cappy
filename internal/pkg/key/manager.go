package key

import (
	"fmt"
	"github.com/go-errors/errors"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	ErrNotDir         = errors.New("path is not a directory")
	ErrInvalidKeySize = errors.New("invalid key size")
)

const fileExt = ".key"

type Manager interface {
	List() ([]string, error)
	Read(id string) (Key, error)
	Write(Key) error
	Delete(id string) error
}

type manager struct {
	path string
}

func NewManager(path string) (Manager, error) {
	info, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("file info failed: %v", err)
	}

	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return nil, fmt.Errorf("failed to create dir: %v", err)
		}
	} else if !info.IsDir() {
		return nil, ErrNotDir
	}

	return &manager{path: path}, nil
}

func (m *manager) List() ([]string, error) {
	l := make([]string, 0)

	err := filepath.Walk(m.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) != fileExt {
			return nil
		}

		dot := strings.LastIndexByte(info.Name(), '.')
		domain := info.Name()[:dot]

		l = append(l, domain)

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("listing failed: %v", err)
	}

	return l, nil
}

func (m *manager) getPath(id string) string {
	return path.Join(m.path, strings.ToLower(id)+fileExt)
}

func (m *manager) Read(id string) (Key, error) {
	file, err := os.OpenFile(m.getPath(id), os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to open key file: %v", err)
	}
	defer file.Close()

	key := Key{}
	if err := key.ReadFrom(file); err != nil {
		return nil, fmt.Errorf("reading key failed: %v", err)
	}

	return key, nil
}

func (m *manager) Write(key Key) error {
	file, err := os.OpenFile(m.getPath(key.Identifier()), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to open key file: %v", err)
	}
	defer file.Close()

	if err := key.WriteTo(file); err != nil {
		return fmt.Errorf("error writing key: %v", err)
	}

	return nil
}

func (m *manager) Delete(id string) error {
	return os.Remove(m.getPath(id))
}
