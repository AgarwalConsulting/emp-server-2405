package repository

import "algogrit.com/emp_server/entities"

type inmemRepo struct {
	employees []entities.Employee
}

func (repo *inmemRepo) ListAll() ([]entities.Employee, error) {
	return repo.employees, nil
}

func (repo *inmemRepo) Save(newEmp entities.Employee) (*entities.Employee, error) {
	newEmp.ID = len(repo.employees) + 1

	repo.employees = append(repo.employees, newEmp)

	return &newEmp, nil
}

func NewInMem() EmployeeRepository {
	var employees = []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
		{2, "Shikhar", "Cloud", 10002},
		{3, "Mark", "SRE", 2003},
	}

	return &inmemRepo{employees: employees}
}
