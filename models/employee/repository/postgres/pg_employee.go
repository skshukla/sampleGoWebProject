package postgres

import (
	models "../../../../models"
	"github.com/jinzhu/gorm"
)
type employeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(DB *gorm.DB) *employeeRepository {
	return &employeeRepository{DB: DB}
}

func (o *employeeRepository) GetEmployeeById(id int) models.Employee {
	var emp = models.Employee{}
	o.DB.Model(models.Employee{}).Where("id = ?", id).Find(&emp)
	return emp
}


func (o *employeeRepository) SaveEmployee(emp *models.Employee) models.Employee {
	o.DB.Save(emp)
	return *emp
}

func (o *employeeRepository) FindAllEmployees() []models.Employee {
	var response []models.Employee
	o.DB.Table("employees").Find(&response)
	return response
}