package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"

	empHTTP "algogrit.com/emp_server/employees/http"
	"algogrit.com/emp_server/employees/service"
	"algogrit.com/emp_server/entities"
)

func TestIndexV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockV1Svc := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.New(mockV1Svc)

	expectedEmp := []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
	}

	mockV1Svc.EXPECT().Index().Return(expectedEmp, nil)

	req := httptest.NewRequest("GET", "/v1/employees", nil)
	resRec := httptest.NewRecorder()

	sut.ServeHTTP(resRec, req)

	assert.Equal(t, http.StatusOK, resRec.Code)

	var actualEmp []entities.Employee

	json.NewDecoder(resRec.Body).Decode(&actualEmp)

	assert.Equal(t, len(expectedEmp), len(actualEmp))
	assert.Equal(t, expectedEmp[0].ID, actualEmp[0].ID)
}

func TestCreateV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockV1Svc := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.New(mockV1Svc)

	expectedEmp := entities.Employee{Name: "Gaurav", Department: "LnD"}
	createdEmp := expectedEmp
	createdEmp.ID = 1

	mockV1Svc.EXPECT().Create(expectedEmp).Return(&createdEmp, nil)

	reqBody := strings.NewReader(`{"name": "Gaurav", "speciality": "LnD"}`)

	req := httptest.NewRequest("POST", "/v1/employees", reqBody)
	resRec := httptest.NewRecorder()

	sut.ServeHTTP(resRec, req)

	assert.Equal(t, http.StatusOK, resRec.Code)

	var actualEmp entities.Employee

	json.NewDecoder(resRec.Body).Decode(&actualEmp)

	assert.Equal(t, createdEmp, actualEmp)
}

func FuzzCreateV1(f *testing.F) {
	f.Add(`{"name": "Gaurav", "speciali`)

	f.Fuzz(func(t *testing.T, jsonString string) {
		ctrl := gomock.NewController(t)
		mockV1Svc := service.NewMockEmployeeService(ctrl)

		sut := empHTTP.New(mockV1Svc)

		expectedEmp := entities.Employee{Name: "Gaurav", Department: "LnD"}
		createdEmp := expectedEmp
		createdEmp.ID = 1

		reqBody := strings.NewReader(jsonString)

		req := httptest.NewRequest("POST", "/v1/employees", reqBody)
		resRec := httptest.NewRecorder()

		sut.ServeHTTP(resRec, req)

		assert.Equal(t, http.StatusBadRequest, resRec.Code)
	})
}
