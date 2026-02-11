package git

import (
	"context"
	"os"
)

type FileAdapter struct{}

func NewFileAdapter() *FileAdapter {
	return &FileAdapter{}
}

func (a *FileAdapter) GetFile(ctx context.Context, repo, path, ref string) ([]byte, error) {
	return os.ReadFile(path)
}

func (a *FileAdapter) CreatePR(ctx context.Context, repo, title, body, head, base string) (string, error) {
	return "https://github.com/mock/pr/1", nil
}
