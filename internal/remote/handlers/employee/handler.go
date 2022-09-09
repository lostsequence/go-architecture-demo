package employee

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yap-arch-demo/internal/domain/dto"

	"github.com/go-chi/chi/v5"
)

type Employee struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Salary float32 `json:"salary"`
}

type Handler struct {
	empStore []Employee
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Failed to get employee", http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, "Failed to get employee", http.StatusInternalServerError)
		return
	}

	var emp Employee
	for _, e := range h.empStore {
		if e.Id == id {
			emp = e
		}
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

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
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

func IncreaseEmployeeSalary(w http.ResponseWriter, r *http.Request) {
	var emp dto.EmployeeDTO

	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		http.Error(w, "Failed to increae employee salary", http.StatusInternalServerError)
		return
	}

	// логика increase salary
	err = increaseEmployeeSalary(emp.Id)

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
