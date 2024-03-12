package main

//https://pkg.go.dev/gonum.org/v1/gonum/mat

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/gonum/matrix/mat64"
)

func train(data dataframe.DataFrame) mat64.Dense {
	numRow, numCol := data.Dims()

	X := mat64.NewDense(numRow, numCol, nil)

	// Set intercept column to 1
	for i := 0; i < numRow; i++ {
		X.Set(i, 0, 1)
	}

	// Fill the rest of the matrix with values from the dataframe
	for i := 0; i < numRow; i++ {
		for j := 0; j < numCol-1; j++ {
			val := data.Elem(i, j).Float()
			X.Set(i, j+1, val) // j+1 because we want to ignore the target column which is the last one in the dataset
		}
	}

	// create target vector as a matrix
	y := mat64.NewDense(numRow, 1, nil)
	for i := 0; i < numRow; i++ {
		val := data.Elem(i, numCol-1).Float() // get the last element in the row
		y.Set(i, 0, val)
	}

	// now do math stuff
	// multiple regression using the normal equations
	// technically ridge regression to avoid univertible matricies
	var Xt mat64.Dense
	Xt.Clone(X.T()) // transpose of the matrix

	// multiply Xt and X to get XtX
	var XtX mat64.Dense
	XtX.Mul(&Xt, X)

	//getting the ridge which is all 0 and 0.01 from top left to bottom right
	reg := mat64.NewDense(numCol, numCol, nil)
	for i := 0; i < numCol; i++ {
		reg.Set(i, i, 0.01)
	}

	// add the ridge to the matrix
	XtX.Add(&XtX, reg)

	// Get the inverse of XtX
	var XtXi mat64.Dense
	err := XtXi.Inverse(&XtX)
	// there is sometimes an error when producing the inverse
	// it should be avoided when using ridge regression
	if err != nil {
		fmt.Println("There was an error:")
		fmt.Println(err)
	}

	// Multiply XtXi with Xt to get XtXiT
	var XtXiT mat64.Dense
	XtXiT.Mul(&XtXi, &Xt)

	// Multiply XtXiT with y to get the coefficients
	var coef mat64.Dense
	coef.Mul(&XtXiT, y)

	return coef
}
