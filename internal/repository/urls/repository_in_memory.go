package urls

import (
	"context"
	"errors"
	"sync"
)

type repositoryInMemory struct {
	data map[string]string
	mu   *sync.RWMutex
}

func NewRepositoryInMemory() *repositoryInMemory {
	return &repositoryInMemory{
		data: make(map[string]string),
		mu:   &sync.RWMutex{},
	}
}

func (r repositoryInMemory) SaveShortenedURL(
	_ context.Context,
	in *ShortenUrlIn,
) (*ShortenURLOut, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.data[in.Alias] == in.OriginalURL {
		return &ShortenURLOut{Alias: in.Alias}, nil
	}

	r.data[in.Alias] = in.OriginalURL
	return &ShortenURLOut{Alias: in.Alias}, nil

}

func (r repositoryInMemory) GetOriginalURL(
	_ context.Context,
	in *GetOriginalURLIn,
) (*GetOriginalURlOut, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if _, ok := r.data[in.ShortURL]; !ok {
		return nil, errors.New("original url not found")
	}

	return &GetOriginalURlOut{OriginalURL: r.data[in.ShortURL]}, nil

}
