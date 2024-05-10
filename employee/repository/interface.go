package repository

import "algogrit.com/emp_server/entities"

type EmployeeRepository interface {
	ListAll() ([]entities.Employee, error)
	Save(entities.Employee) (*entities.Employee, error)
}
