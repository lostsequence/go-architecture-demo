package employee

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yap-arch-demo/internal/domain/dto"

	"github.com/go-chi/chi/v5"
)

type EmployeeService interface {
	Get(id int) (*dto.EmployeeDTO, error)
	IncreaseEmployeeSalary(id int) error
	CreateEmployee(emp dto.EmployeeDTO) error
}

type EmployeeHandler struct {
	employeeService EmployeeService
}

func NewEmployeeHandler(es EmployeeService) EmployeeHandler {
	return EmployeeHandler{
		employeeService: es,
	}
}

func (eh EmployeeHandler) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Failed to get employee", http.StatusInternalServerError)
		return
	}

	emp, err := eh.employeeService.Get(id)

	if err != nil {
		http.Error(w, "Failed to get employee", http.StatusInternalServerError)
		return
	}

	empJson, err := json.Marshal(emp)

	if err != nil {
		http.Error(w, "Failed to get employee", http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(empJson); err != nil {
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
	}
}

func (eh EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp dto.EmployeeDTO

	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
		return
	}

	err = eh.employeeService.CreateEmployee(emp)

	if err != nil {
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
		return
	}

	if _, err = w.Write([]byte("Employee created")); err != nil {
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
	}
}

func (eh EmployeeHandler) IncreaseEmployeeSalary(w http.ResponseWriter, r *http.Request) {
	var emp dto.EmployeeDTO

	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		http.Error(w, "Failed to increae employee salary", http.StatusInternalServerError)
		return
	}

	err = eh.employeeService.IncreaseEmployeeSalary(emp.Id)

	if err != nil {
		http.Error(w, "Failed to increae employee salary", http.StatusInternalServerError)
		return
	}

	newEmp, err := eh.employeeService.Get(emp.Id)

	if err != nil {
		http.Error(w, "Failed to get employee", http.StatusInternalServerError)
		return
	}

	empJson, err := json.Marshal(newEmp)

	if err != nil {
		http.Error(w, "Failed to increae employee salary", http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(empJson); err != nil {
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
	}
}
