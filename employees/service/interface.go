package service

import "algogrit.com/emp_server/entities"

//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOPACKAGE.go

type EmployeeService interface {
	Index() ([]entities.Employee, error)
	Create(newEmp entities.Employee) (*entities.Employee, error)
}
