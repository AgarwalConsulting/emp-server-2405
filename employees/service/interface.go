package service

import "algogrit.com/emp_server/entities"

type EmployeeService interface {
	Index() ([]entities.Employee, error)
	Create(newEmp entities.Employee) (*entities.Employee, error)
}
