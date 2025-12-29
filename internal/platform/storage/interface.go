package storage

import (
	"context"
	"io"
)

type Service interface {
	Upload(ctx context.Context, file io.Reader, filename string) (string, error)
	GetURL(filename string) string
}
