package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type LocalStorage struct {
	BaseDir string
	BaseURL string
}

func NewLocalStorage(baseDir, baseURL string) *LocalStorage {
	// Ensure directory exists
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create storage directory: %v", err))
	}
	return &LocalStorage{
		BaseDir: baseDir,
		BaseURL: baseURL,
	}
}

func (s *LocalStorage) Upload(ctx context.Context, file io.Reader, filename string) (string, error) {
	// Generate unique filename to avoid collisions
	ext := filepath.Ext(filename)
	name := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	path := filepath.Join(s.BaseDir, name)

	// Create file
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy content
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return s.GetURL(name), nil
}

func (s *LocalStorage) GetURL(filename string) string {
	return fmt.Sprintf("%s/%s", s.BaseURL, filename)
}
