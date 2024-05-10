package repository

import "algogrit.com/emp_server/entities"

//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOPACKAGE.go

type EmployeeRepository interface {
	ListAll() ([]entities.Employee, error)
	Save(entities.Employee) (*entities.Employee, error)
}
