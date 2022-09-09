package remote

import (
	"yap-arch-demo/internal/remote/handlers/employee"

	"github.com/go-chi/chi/v5"
)

func NewRouter(emp employee.EmployeeHandler) chi.Router {
	r := chi.NewRouter()
	r.Get("/{id}", emp.Get)
	r.Post("/create", emp.CreateEmployee)
	r.Post("/addsalary", emp.IncreaseEmployeeSalary)
	return r
}
