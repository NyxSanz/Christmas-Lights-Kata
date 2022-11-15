package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMatrix(t *testing.T) {

	expected := [][]int{{0, 0}, {0, 0}}
	got := createMatrix(2, 2)

	assert.Equal(t, expected, got)

}

func TestCheckIfAllLightTurnOff(t *testing.T) {

	expected := false

	got := checkIfAllLighTurnOff()

	assert.Equal(t, expected, got, "all light are turn off")

}

func TestCheckifAllLightTurnOffFailed(t *testing.T) {

	expected := true
	got := checkIfAllLighTurnOff()

	assert.NotEqual(t, expected, got, "all light are not turn off")
}

func TestCheckAllLightTurnOn(t *testing.T) {
	expected := true

	matrix := createMatrix(10, 10)
	matrix = turnOnAllLight(matrix)
	got := checkIfAllLightTurnOn(matrix)

	assert.Equal(t, expected, got)
}

func TestTurnOnLigh(t *testing.T) {

	expected := 1000000

	matrix := createMatrix(1000, 1000)

	var positionX, positionY position

	positionX.start = 0
	positionX.end = 999
	positionY.start = 0
	positionY.end = 999

	matrix = turnOnLight(positionX, positionY, matrix)

	got := countLight(matrix)

	assert.Equal(t, expected, got)
}

func TestTurnOffLight(t *testing.T) {
	expected := 999996

	matrix := createMatrix(1000, 1000)
	matrix = turnOnAllLight(matrix)

	var positionX, positionY position

	positionX.start = 499
	positionX.end = 500
	positionY.start = 499
	positionY.end = 500

	matrix = turnOffLight(positionX, positionY, matrix)

	got := countLight(matrix)

	assert.Equal(t, expected, got)
}

func TestToggleLightOffToOn(t *testing.T) {

	expected := 1000
	matrix := createMatrix(1000, 1000)

	var positionX, positionY position

	positionX.start = 0
	positionX.end = 999
	positionY.start = 0
	positionY.end = 0

	matrix = toggleLight(positionX, positionY, matrix)

	got := countLight(matrix)

	assert.Equal(t, expected, got)

}

func TestToggleLightOnToOff(t *testing.T) {

	expected := 999000
	matrix := createMatrix(1000, 1000)
	matrix = turnOnAllLight(matrix)
	var positionX, positionY position

	positionX.start = 0
	positionX.end = 999
	positionY.start = 0
	positionY.end = 0

	matrix = toggleLight(positionX, positionY, matrix)

	got := countLight(matrix)

	assert.Equal(t, expected, got)

}
