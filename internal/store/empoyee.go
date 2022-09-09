package store

import (
	"sync"
	"yap-arch-demo/internal/domain/entity"
	"yap-arch-demo/internal/domain/errors"
)

type EmployeeStore struct {
	inMemory []*entity.EmployeeEntity
	mu       sync.Mutex
}

func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{
		inMemory: make([]*entity.EmployeeEntity, 0),
		mu:       sync.Mutex{},
	}
}

func (es *EmployeeStore) Create(employee *entity.EmployeeEntity) error {
	es.mu.Lock()
	defer es.mu.Unlock()

	es.inMemory = append(es.inMemory, employee)
	return nil
}

func (es *EmployeeStore) Get(id int) (*entity.EmployeeEntity, error) {
	es.mu.Lock()
	defer es.mu.Unlock()

	for _, e := range es.inMemory {
		if e.Id == id {
			return e, nil
		}
	}

	return nil, errors.NotFound
}

func (es *EmployeeStore) Update(employee *entity.EmployeeEntity) error {
	emp, err := es.Get(employee.Id)

	if err != nil {
		return err
	}

	emp.Name = employee.Name
	emp.Salary = employee.Salary
	emp.SecurityCardId = employee.SecurityCardId

	return nil
}
