package main

import (
	"context"
	"fmt"
	"sampleGoWebProject/container"
	"sampleGoWebProject/models"
	"testing"
)

func TestGetEmployeeById(t *testing.T)  {

	c := &container.Container{Profile: "test"}
	c.WireDependencies()

	e1 := &models.Employee{Name: "Abcd"}
	savedEmp := c.GetEmployeeUseCase().SaveEmployee(context.Background(), e1)

	if savedEmp.Id == 0 {
		t.Errorf("Saved Employee Id cannot be zero")
	}
	e2 := c.GetEmployeeUseCase().GetEmployeeById(context.Background(), savedEmp.Id)
	if e2.Name != e1.Name {
		t.Errorf("Expected name {%s}, found name {%s}", e1.Name, e2.Name)
	}
	fmt.Println(fmt.Sprintf("Count of total employees is {%d}", len(c.GetEmployeeUseCase().FindAllEmployees())))

}
