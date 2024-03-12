package main

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/gonum/matrix/mat64"
)

func test(coef mat64.Dense, testData dataframe.DataFrame) *mat64.Dense {

	// get dimensions of testdata
	numRow, numCol := testData.Dims()

	// create a new matrix with 0 in every position
	X := mat64.NewDense(numRow, numCol, nil)

	// fill the matrix
	for i := 0; i < numRow; i++ {
		for j := 0; j < numCol-1; j++ {
			val := testData.Elem(i, j).Float()
			X.Set(i, j+1, val)
		}
		X.Set(i, 0, 1) // Set intercept column to 1
	}

	// create prediction matrix that is a duplicate of X
	var yPredActual mat64.Dense
	yPredActual.Apply(func(i int, j int, v float64) float64 {
		return v
	}, X)

	// Multiply X with coefficients to get predictions
	var predictions mat64.Dense
	predictions.Mul(X, &coef)

	// add the predicted values to the yPredActual matrix
	yPredActual.Apply(func(i int, j int, v float64) float64 {
		return predictions.At(i, 0)
	}, &yPredActual)

	// add actual values to the yPredActual matrix
	for i := 0; i < numRow; i++ {
		yPredActual.Set(i, numCol-1, testData.Elem(i, numCol-1).Float())
	}

	return &yPredActual
}
