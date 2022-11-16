package main

import (
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tabaproc/testify-study/ex3/data"
	"github.com/tabaproc/testify-study/ex3/mocks"
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
	repoMock := new(mocks.RepoMock)
	staff := data.Staff{
		ID:   r.staffID,
		Name: data.StaffData[r.staffID]["name"].(string),
		Post: data.StaffData[r.staffID]["post"].(string),
		Rate: data.StaffData[r.staffID]["rate"].(float64),
	}
	staffSecond := staff
	staffSecond.Rate += data.UpRate[data.StaffData[r.staffID]["post"].(string)]

	repoMock.On("Staff", r.staffID).Return(staff)
	repoMock.On("UpdateIncome", &staff).Times(1)
	repoMock.On("UpdateIncome", &staffSecond).Times(1)

	respStaff := execute(r.staffID, repoMock)

	expected := staff.Rate + (data.UpRate[staff.Post] * 2)
	assert.Equal(r.T(), expected, respStaff.Rate)
}

// execute()のテスト Mockの引数が正しくない
func (r *MainSuite) TestExecuteError1() {
	repoMock := new(mocks.RepoMock)
	staff := data.Staff{
		ID:   r.staffID,
		Name: data.StaffData[r.staffID]["name"].(string),
		Post: data.StaffData[r.staffID]["post"].(string),
		Rate: data.StaffData[r.staffID]["rate"].(float64),
	}
	staffSecond := staff
	staffSecond.Rate += data.UpRate[data.StaffData[r.staffID]["post"].(string)]

	repoMock.On("Staff", 2).Return(staff)
	repoMock.On("UpdateIncome", &staff).Times(1)
	repoMock.On("UpdateIncome", &staffSecond).Times(1)

	respStaff := execute(r.staffID, repoMock)

	expected := staff.Rate + (data.UpRate[staff.Post] * 2)
	assert.Equal(r.T(), expected, respStaff.Rate)
}

// execute()のテスト MockでAnythingを使う
func (r *MainSuite) TestExecuteAnything() {
	repoMock := new(mocks.RepoMock)
	staff := data.Staff{
		ID:   r.staffID,
		Name: data.StaffData[r.staffID]["name"].(string),
		Post: data.StaffData[r.staffID]["post"].(string),
		Rate: data.StaffData[r.staffID]["rate"].(float64),
	}
	staffSecond := staff
	staffSecond.Rate += data.UpRate[data.StaffData[r.staffID]["post"].(string)]

	repoMock.On("Staff", mock.AnythingOfType("int")).Return(staff)
	repoMock.On("UpdateIncome", &staff).Times(1)
	repoMock.On("UpdateIncome", &staffSecond).Times(1)

	respStaff := execute(r.staffID, repoMock)

	expected := staff.Rate + (data.UpRate[staff.Post] * 2)
	assert.Equal(r.T(), expected, respStaff.Rate)
}

// executeのテスト Returnが正しくない
func (r *MainSuite) TestExecuteError2() {
	repoMock := new(mocks.RepoMock)
	staff := data.Staff{
		ID:   r.staffID,
		Name: "",
		Post: "",
	}
	staffSecond := staff
	staffSecond.Rate += data.UpRate["Member"]

	repoMock.On("Staff", r.staffID).Return(staff)
	repoMock.On("UpdateIncome", &staff).Times(1)
	repoMock.On("UpdateIncome", &staffSecond).Times(1)

	respStaff := execute(r.staffID, repoMock)

	expected := staff.Rate + (data.UpRate[staff.Post] * 2)
	assert.Equal(r.T(), expected, respStaff.Rate)
}
