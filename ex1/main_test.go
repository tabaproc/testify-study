package main

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"
)

// testingでsquare関数をテスト 正常
func TestTestingSquareNormal(t *testing.T) {
	expected := 25
	if expected != square(5) {
		t.Error("二乗した結果と想定したの値がマッチしません")
	}
}

// testingでsquare関数をテスト エラー
func TestTestingSquareError(t *testing.T) {
	expected := 24
	if expected != square(5) {
		t.Error("二乗した結果と想定したの値がマッチしません")
	}
}

// testifyでsquare関数をテスト 正常
func TestTestifySquareNormal(t *testing.T) {
	expected := 25
	assert.Equal(t, expected, square(5), "正常の場合、このメッセージは表示されない")
}

// testifyでsquare関数をテスト エラー
func TestTestifySquareError1(t *testing.T) {
	expected := 24
	assert.Equal(t, expected, square(5), "予測の値と実際の値が違う")
	expected = 23
	assert.Equal(t, expected, square(5), "予測の値と実際の値が違う")
}

// testifyでsquare関数をテスト エラー
func TestTestifySquareError2(t *testing.T) {
	expected := 24
	require.Equal(t, expected, square(5), "予測の値と実際の値が違う")
	expected = 23
	require.Equal(t, expected, square(5), "予測の値と実際の値が違う")
}
