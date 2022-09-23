package service

import (
	"errors"
	swagger "github.com/xpoh/BRE-test/cmd/go"
)

type Storager interface {
	addToStorage(content swagger.Content) error
	readFromStorage(id string) (string, error)
	removeFromStorage(id string) error
}

type Service struct {
	pipe         map[string]string
	cacheStorage Storager
	dbStorage    Storager
}

func NewService(cacheStorage Storager, dbStorage Storager) *Service {
	s := &Service{cacheStorage: cacheStorage, dbStorage: dbStorage}
	s.pipe = make(map[string]string, 0)
	return s
}

func (s *Service) saveContent(c swagger.Content) error {
	if _, ok := s.pipe[c.Id]; ok {
		return errors.New("Internal error")
	}
	s.pipe[c.Id] = c.Body

	err := s.cacheStorage.addToStorage(c)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) readFromStorage(id string) (string, error) {
	if _, ok := s.pipe[id]; ok {
		return "", errors.New("not find")
	}
	return s.cacheStorage.readFromStorage(id)
}
