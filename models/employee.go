package models

import (
	"context"
)

type Employee struct {
	Id int
	Name string
}

type EmployeeUseCase interface {
	GetEmployeeById(context context.Context, id int) Employee
	SaveEmployee(context context.Context, emp *Employee) Employee
	FindAllEmployees() []Employee
}

type EmployeeRepository interface {
	GetEmployeeById(id int) Employee
	SaveEmployee(emp *Employee) Employee
	FindAllEmployees() []Employee

}

