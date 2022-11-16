package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// testifyでsuiteを使わずsquare関数をテスト 正常
func TestNoSuiteSquareNormal(t *testing.T) {
	expected := 25
	assert.Equal(t, expected, square(5))
}

// Suiteを使う
type MainSuite struct {
	suite.Suite
}

func TestMainSuite(t *testing.T) {
	suite.Run(t, new(MainSuite))
}

func (r *MainSuite) TestSuiteSquareNormal() {
	expected := 25
	assert.Equal(r.T(), expected, square(5))
}

func (r *MainSuite) TestSuiteSquareError() {
	expected := 24
	assert.Equal(r.T(), expected, square(5))
}

// 構造体生成時にCall
func (r *MainSuite) SetupSuite() {
	r.T().Log("SetupSuite")
}

// テストメソッド呼び出し時にCall
func (r *MainSuite) SetupTest() {
	r.T().Log("SetupTest")
}

// テストメソッド呼び出し時にCall
func (r *MainSuite) BeforeTest(suiteName, testName string) {
	r.T().Logf("BeforeTest: %s.%s", suiteName, testName)
}
