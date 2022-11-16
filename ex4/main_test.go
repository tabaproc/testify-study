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
	repoMock *mocks.RepositoryOperator
	staffID  int
}

// TestMainSuite
func TestMainSuite(t *testing.T) {
	suite.Run(t, new(MainSuite))
}

// テストメソッド呼び出し時にCall
func (r *MainSuite) SetupTest() {
	r.repoMock = mocks.NewRepositoryOperator(r.T())
	r.staffID = 4
}

// execute()のテスト 正常
func (r *MainSuite) TestExecuteNormal() {
	staff := data.Staff{
		ID:   r.staffID,
		Name: data.StaffData[r.staffID]["name"].(string),
		Post: data.StaffData[r.staffID]["post"].(string),
		Rate: data.StaffData[r.staffID]["rate"].(float64),
	}

	r.repoMock.On("Staff", r.staffID).Return(staff)
	r.repoMock.On("UpdateIncome", &staff)

	respStaff := execute(r.staffID, r.repoMock)

	expected := staff.Rate
	assert.Equal(r.T(), expected, respStaff.Rate)
}

// execute()のテスト Mockの引数が正しくない
func (r *MainSuite) TestExecuteError1() {
	staff := data.Staff{
		ID:   r.staffID,
		Name: data.StaffData[r.staffID]["name"].(string),
		Post: data.StaffData[r.staffID]["post"].(string),
		Rate: data.StaffData[r.staffID]["rate"].(float64),
	}

	r.repoMock.On("Staff", 2).Return(staff)
	r.repoMock.On("UpdateIncome", &staff)

	respStaff := execute(r.staffID, r.repoMock)

	expected := staff.Rate
	assert.Equal(r.T(), expected, respStaff.Rate)
}

// execute()のテスト MockでAnythingを使う
func (r *MainSuite) TestExecuteAnything() {
	staff := data.Staff{
		ID:   r.staffID,
		Name: data.StaffData[r.staffID]["name"].(string),
		Post: data.StaffData[r.staffID]["post"].(string),
		Rate: data.StaffData[r.staffID]["rate"].(float64),
	}

	r.repoMock.On("Staff", mock.AnythingOfType("int")).Return(staff)
	r.repoMock.On("UpdateIncome", &staff)

	respStaff := execute(r.staffID, r.repoMock)

	expected := staff.Rate
	assert.Equal(r.T(), expected, respStaff.Rate)
}
