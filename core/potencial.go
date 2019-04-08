package core
func Potencial(cost [][]float64, startSolution [][]float64, iteration int) [][]float64 {
	numIteration:=10
	finalSolution := startSolution
	if iteration > numIteration{
		return finalSolution
	}
	iteration = iteration+1
	u := make([]float64, len(cost))
	v := make([]float64, len(cost[0]))
	for i := 0; i < len(cost); i++ {
		for j := 0; j < len(cost[0]); j++ {
			if isExist(u[i]) {
				v[j] = cost[i][j]
			} else if isExist(v[j]) {
				u[i] = cost[i][j] - v[j]
			}
		}
	}
	if isAllExist(v) && isAllExist(u[1:]) {
		delta, isOptimal := computeDelta(u, v, cost, startSolution)
		if !isOptimal {
			finalSolution = redistribution(delta, cost, finalSolution)
			finalSolution = Potencial(cost, finalSolution, iteration)
		}
	} else {finalSolution = Potencial(cost, finalSolution, iteration)} 
	return finalSolution
}
func computeDelta(u, v []float64, cost, northWest [][]float64) ([][]float64, bool) {
	var delta [][]float64
	isOptimal := true
	for i := 0; i < len(cost); i++ {
		for j := 0; j < len(cost[0]); j++ {
			if northWest[i][j] != 0 {
				delta[i][j] = 0
			} else {
				delta[i][j] = cost[i][j] - u[i] - v[j]
				if delta[i][j] < 0 {
					isOptimal = false
				}
			}
		}
	}
	return delta, isOptimal
}
func redistribution(delta, cost, startSolution [][]float64) [][]float64 {
	finishSolution := startSolution
	i, j, _ := minElem(delta)
	r, c, min := minNeighbor(startSolution, i, j)
	finishSolution[r][c] = min
	if finishSolution[r+1][c] != 0 {
		finishSolution[r+1][c] = finishSolution[r+1][c] - min
	} else if finishSolution[r-1][c] != 0 {
		finishSolution[r-1][c] = finishSolution[r-1][c] - min
	} else if finishSolution[r][c-1] != 0 {
		finishSolution[r][c-1] = finishSolution[r][c-1] - min
	} else if finishSolution[r][c+1] != 0 {
		finishSolution[r][c+1] = finishSolution[r][c+1] - min
	} else if finishSolution[r-1][c-1] != 0 {
		finishSolution[r-1][c-1] = finishSolution[r-1][c-1] + min
	} else if finishSolution[r-1][c+1] != 0 {
		finishSolution[r-1][c+1] = finishSolution[r-1][c+1] + min
	} else if finishSolution[r+1][c+1] != 0 {
		finishSolution[r+1][c+1] = finishSolution[r+1][c+1] + min
	} else if finishSolution[r+1][c-1] != 0 {
		finishSolution[r+1][c-1] = finishSolution[r+1][c-1] + min
	}
	return finishSolution
}
