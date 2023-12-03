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

	if _, ok := r.data[in.OriginalURL]; ok {
		return &ShortenURLOut{Alias: r.data[in.OriginalURL]}, nil
	}

	r.data[in.OriginalURL] = in.Alias
	return &ShortenURLOut{Alias: r.data[in.OriginalURL]}, nil

}

func (r repositoryInMemory) GetOriginalURL(
	_ context.Context,
	in *GetOriginalURLIn,
) (*GetOriginalURlOut, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for url, alias := range r.data {
		if alias == in.ShortURL {
			return &GetOriginalURlOut{OriginalURL: url}, nil
		}
	}

	return nil, errors.New("original url not found")
}
