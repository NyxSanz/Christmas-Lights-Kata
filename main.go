package main

import "fmt"

type position struct {
	start, end int
}

func main() {
	//fmt.Print("test")
	fmt.Println(createMatrix(2, 2))

}

func createMatrix(row, column int) [][]int {

	m := make([][]int, row)
	for i := range m {
		m[i] = make([]int, column)
	}
	return m

}

func checkIfAllLighTurnOff(matrix [][]int) bool {

	return countLight(matrix) > 0

}

func countLight(matrix [][]int) int {

	total := 0

	for i, row := range matrix {

		for j := range row {

			total = total + matrix[i][j]

		}
	}

	return total
}

func turnOnAllLight(matrix [][]int) [][]int {

	for i, row := range matrix {
		for j := range row {
			matrix[i][j] = 1
		}
	}

	return matrix

}

func checkIfAllLightTurnOn(matrix [][]int) bool {

	for i, row := range matrix {
		for j := range row {
			if matrix[i][j] == 0 {
				return false
			}

		}
	}

	return true
}

func turnOnLight(postionX, positionY position, matrix [][]int) [][]int {

	for i := postionX.start; i <= postionX.end; i++ {
		for j := positionY.start; j <= positionY.end; j++ {
			matrix[i][j]++
		}
	}

	return matrix

}

func turnOffLight(postionX, positionY position, matrix [][]int) [][]int {

	for i := postionX.start; i <= postionX.end; i++ {
		for j := positionY.start; j <= positionY.end; j++ {

			if matrix[i][j] != 0 {
				matrix[i][j] = 0
			}

		}
	}

	return matrix

}

func toggleLight(postionX, positionY position, matrix [][]int) [][]int {

	for i := postionX.start; i <= postionX.end; i++ {
		for j := positionY.start; j <= positionY.end; j++ {

			if matrix[i][j] == 0 {
				matrix[i][j]++
				continue
			}
			matrix[i][j] = 0

		}
	}

	return matrix

}
