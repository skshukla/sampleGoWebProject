package usecase

import (
	"context"
	"sampleGoWebProject/models"
)

type employeeUseCase struct {
	employeeRepository models.EmployeeRepository
}

func NewEmployeeUseCase(employeeRepository models.EmployeeRepository)  *employeeUseCase {
	return &employeeUseCase{employeeRepository}
}

func (o *employeeUseCase)GetEmployeeById(context context.Context, id int) models.Employee{
	return o.employeeRepository.GetEmployeeById(id)
}


func (o *employeeUseCase)SaveEmployee(context context.Context, emp *models.Employee) models.Employee{
	return o.employeeRepository.SaveEmployee(emp)
}

func (o *employeeUseCase) FindAllEmployees() []models.Employee {
	return o.employeeRepository.FindAllEmployees()
}

