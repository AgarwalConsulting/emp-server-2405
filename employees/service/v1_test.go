package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"

	"algogrit.com/emp_server/employees/repository"
	"algogrit.com/emp_server/employees/service"
	"algogrit.com/emp_server/entities"
)

func TestIndex(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := repository.NewMockEmployeeRepository(ctrl)

	sut := service.NewV1(mockRepo)

	expectedEmp := []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
	}

	mockRepo.EXPECT().ListAll().Return(expectedEmp, nil)

	actualEmp, err := sut.Index()

	assert.Nil(t, err)
	assert.NotNil(t, actualEmp)
	assert.Equal(t, expectedEmp, actualEmp)
}
