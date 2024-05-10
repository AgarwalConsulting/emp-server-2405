package repository

import (
	"sync"

	"algogrit.com/emp_server/entities"
)

type inmemRepo struct {
	employees []entities.Employee
	sync.RWMutex
}

func (repo *inmemRepo) ListAll() ([]entities.Employee, error) {
	repo.RWMutex.RLock()
	defer repo.RWMutex.RUnlock()
	return repo.employees, nil
}

func (repo *inmemRepo) Save(newEmp entities.Employee) (*entities.Employee, error) {
	repo.RWMutex.Lock()
	defer repo.RWMutex.Unlock()

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
