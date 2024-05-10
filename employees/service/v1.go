package service

import (
	"algogrit.com/emp_server/employees/repository"
	"algogrit.com/emp_server/entities"
)

type empV1Svc struct {
	repo repository.EmployeeRepository
}

func (svc empV1Svc) Index() ([]entities.Employee, error) {
	return svc.repo.ListAll()
}

func (svc empV1Svc) Create(newEmp entities.Employee) (*entities.Employee, error) {
	return svc.repo.Save(newEmp)
}

func NewV1(repo repository.EmployeeRepository) EmployeeService {
	return empV1Svc{repo}
}
