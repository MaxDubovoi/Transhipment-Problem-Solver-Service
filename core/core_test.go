package core

import (
	"testing"
)


func TestNorthWest(t *testing.T) {

	providers := []float64{30, 40, 20}

	consumers := []float64{20, 30, 30, 10}
	
	var resultsAssert [][]float64
	resultsAssert = append(resultsAssert, []float64{20,10,0,0})
	resultsAssert = append(resultsAssert, []float64{0,20,20,0})
	resultsAssert = append(resultsAssert, []float64{0,0,10,10})



	got := NorthWest(providers, consumers)
	if got[0][0]!=resultsAssert[0][0]{
		t.Errorf("Value = %v; Want= %v", got[0][0], resultsAssert[0][0])
	}else if got[0][1]!=resultsAssert[0][1]{
		t.Errorf("Value = %v; Want= %v", got[0][0], resultsAssert[0][0])
	}	
}
