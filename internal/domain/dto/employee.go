package dto

type EmployeeDTO struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Salary float32 `json:"salary"`
}
