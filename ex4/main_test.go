package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/tabaproc/testify-study/ex4/data"
	"github.com/tabaproc/testify-study/ex4/mocks"
)

// MainSuite
type MainSuite struct {
	suite.Suite
	staffID int
}

// TestMainSuite
func TestMainSuite(t *testing.T) {
	suite.Run(t, new(MainSuite))
}

// テストメソッド呼び出し時にCall
func (r *MainSuite) SetupSuite() {
	r.staffID = 4
}

// execute()のテスト 正常
func (r *MainSuite) TestExecuteNormal() {
	repoMock := mocks.NewRepositoryOperator(r.T())
	staff := data.Staff{
		ID:   r.staffID,
		Name: data.StaffData[r.staffID]["name"].(string),
		Post: data.StaffData[r.staffID]["post"].(string),
		Rate: data.StaffData[r.staffID]["rate"].(float64),
	}

	repoMock.On("Staff", r.staffID).Return(staff)
	repoMock.On("UpdateIncome", &staff)

	respStaff := execute(r.staffID, repoMock)

	expected := staff.Rate
	assert.Equal(r.T(), expected, respStaff.Rate)
}

// execute()のテスト Mockの引数が正しくない
func (r *MainSuite) TestExecuteError1() {
	repoMock := mocks.NewRepositoryOperator(r.T())
	staff := data.Staff{
		ID:   r.staffID,
		Name: data.StaffData[r.staffID]["name"].(string),
		Post: data.StaffData[r.staffID]["post"].(string),
		Rate: data.StaffData[r.staffID]["rate"].(float64),
	}

	repoMock.On("Staff", 2).Return(staff)
	repoMock.On("UpdateIncome", &staff)

	respStaff := execute(r.staffID, repoMock)

	expected := staff.Rate
	assert.Equal(r.T(), expected, respStaff.Rate)
}

// execute()のテスト MockでAnythingを使う
func (r *MainSuite) TestExecuteAnything() {
	repoMock := mocks.NewRepositoryOperator(r.T())
	staff := data.Staff{
		ID:   r.staffID,
		Name: data.StaffData[r.staffID]["name"].(string),
		Post: data.StaffData[r.staffID]["post"].(string),
		Rate: data.StaffData[r.staffID]["rate"].(float64),
	}

	repoMock.On("Staff", mock.AnythingOfType("int")).Return(staff)
	repoMock.On("UpdateIncome", &staff)

	respStaff := execute(r.staffID, repoMock)

	expected := staff.Rate
	assert.Equal(r.T(), expected, respStaff.Rate)
}
