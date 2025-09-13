package storage

import "errors"

var ErrLinkNotFound = errors.New("link not found")

type MemoryStorage struct {
	links map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		links: make(map[string]string),
	}
}

func (s *MemoryStorage) Get(code string) (string, error) {
	url, found := s.links[code]
	if !found {
		return "", ErrLinkNotFound
	}

	return url, nil
}

func (s *MemoryStorage) Put(code, url string) error {
	s.links[code] = url
	return nil
}