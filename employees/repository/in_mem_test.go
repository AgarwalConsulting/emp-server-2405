package repository_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"algogrit.com/emp_server/employees/repository"
	"algogrit.com/emp_server/entities"
)

func TestConsistency(t *testing.T) {
	sut := repository.NewInMem()

	existingEmps, err := sut.ListAll()

	assert.Nil(t, err)
	assert.NotNil(t, existingEmps)
	assert.Equal(t, 3, len(existingEmps))

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			newEmp := entities.Employee{1, "Gaurav", "LnD", 1001}
			sut.Save(newEmp)
			sut.ListAll()
		}()
	}

	wg.Wait()

	actualEmps, err := sut.ListAll()

	assert.Nil(t, err)
	assert.NotNil(t, actualEmps)
	assert.Equal(t, 103, len(actualEmps))
}
