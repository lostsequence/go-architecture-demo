package employee

import (
	"errors"
	"log"
	"yap-arch-demo/internal/domain/dto"
	"yap-arch-demo/internal/domain/entity"
	serr "yap-arch-demo/internal/domain/errors"
)

type EmployeeStore interface {
	Create(employee *entity.EmployeeEntity) error
	Get(id int) (*entity.EmployeeEntity, error)
	Update(employee *entity.EmployeeEntity) error
}

type Notifier interface {
	Notify(message string)
}

type EmployeeService struct {
	notifier      Notifier
	employeeStore EmployeeStore
}

func NewEmployeeService(es EmployeeStore) EmployeeService {
	return EmployeeService{
		employeeStore: es,
	}
}

func (s EmployeeService) Get(id int) (*dto.EmployeeDTO, error) {
	emp, err := s.employeeStore.Get(id)

	if err != nil {
		return nil, err
	}

	empDto := &dto.EmployeeDTO{
		Id:     emp.Id,
		Name:   emp.Name,
		Salary: emp.Salary,
	}

	return empDto, nil
}

func (s EmployeeService) CreateEmployee(emp dto.EmployeeDTO) error {
	empEntity := &entity.EmployeeEntity{
		Id:             emp.Id,
		Name:           emp.Name,
		Salary:         emp.Salary,
		SecurityCardId: 123456, // TODO: implement SercurityCardId generation logic
	}

	err := s.employeeStore.Create(empEntity)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s EmployeeService) IncreaseEmployeeSalary(id int) error {
	emp, err := s.employeeStore.Get(id)

	if errors.Is(err, serr.NotFound) {
		return err
	}

	emp.Salary = calculateNewSalary(emp.Salary)

	err = s.employeeStore.Update(emp)

	if err != nil {
		return err
	}

	//TODO: implement employee notification
	s.notifier.Notify("Поздравляем")

	return nil
}
