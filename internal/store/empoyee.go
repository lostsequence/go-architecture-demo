package store

import (
	"yap-arch-demo/internal/domain/entity"
	"yap-arch-demo/internal/domain/errors"
)

type Store struct {
	empStore []*entity.EmployeeEntity
}

func (s Store) Get(id int) (*entity.EmployeeEntity, error) {
	for _, e := range s.empStore {
		if e.Id == id {
			return e, nil
		}
	}

	return nil, errors.NotFound
}
