package core

func isExist(val float64) bool { //return true if val is real
	return val == 0
}
func isAllExist(arr []float64) bool {
	for _, val := range arr {
		if !isExist(val) {
			return false
		}
	}
	return true
}
func minElem(arr [][]float64) (int, int, float64) {
	min := arr[0][0]
	resultRowIndex := 0
	resultColumnIndex := 0

	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr[0]); j++ {
			if min > arr[i][j] {
				min = arr[i][j]
				resultRowIndex = i
				resultColumnIndex = j
			}

		}
	}
	return resultRowIndex, resultColumnIndex, min
}
func minNeighbor(solution [][]float64, rowID, colID int) (int, int, float64) {
	min := solution[rowID][colID]
	resultRow:=rowID
	resultCol := colID
	if min > solution[rowID][colID+1]&&solution[rowID][colID+1]!=0  {
		min = solution[rowID][colID+1]
		resultCol = colID + 1
	} else if min > solution[rowID][colID-1]&&solution[rowID][colID-1]!=0 {
		min = solution[rowID][colID+1]
		resultCol = colID - 1
	} else if min > solution[rowID+1][colID]&&solution[rowID+1][colID]!=0{
		min = solution[rowID+1][colID]
		resultRow = rowID+1
	} else if min > solution[rowID-1][colID]&& solution[rowID-1][colID]!=0{
		min = solution[rowID-1][colID]
		resultRow = rowID-1
	}
	return resultRow, resultCol, min
}
