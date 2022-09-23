package service

import (
	"errors"
)

var App *Service

type Storager interface {
	AddToStorage(content Content) error
	ReadFromStorage(id string) (string, error)
	RemoveFromStorage(id string) error
}

type Service struct {
	pipe         map[string]struct{}
	cacheStorage Storager
	dbStorage    Storager
}

func NewService(cacheStorage Storager, dbStorage Storager) *Service {
	s := &Service{cacheStorage: cacheStorage, dbStorage: dbStorage}
	s.pipe = make(map[string]struct{}, 0)
	return s
}

func (s *Service) SaveContent(c Content) error {
	if _, ok := s.pipe[c.Id]; ok {
		return errors.New("internal error")
	}
	s.pipe[c.Id] = struct{}{}
	err1 := s.dbStorage.AddToStorage(c)
	err2 := s.cacheStorage.AddToStorage(c)
	if err1 == nil && err2 == nil {
		delete(s.pipe, c.Id)
		return nil
	} else {
		return errors.New("error add content")
	}

	return nil
}

func (s *Service) ReadFromStorage(id string) (string, error) {
	if _, ok := s.pipe[id]; ok {
		return "", errors.New("not find")
	}
	return s.cacheStorage.ReadFromStorage(id)
}
