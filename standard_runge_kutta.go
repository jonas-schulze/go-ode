package ode

import "gonum.org/v1/gonum/mat"

func StandardRungeKutta() Solver {
	A := mat.NewBandDense(4, 4, 1, 0, []float64{
		-1, 0,
		0.5, 0,
		0.5, 0,
		1, 0,
	})
	b := mat.NewVecDense(4, []float64{1 / 6, 1 / 3, 1 / 3, 1 / 6})
	c := mat.NewVecDense(4, []float64{0, 0.5, 0.5, 1})
	return ExplicitRungeKutta(A, b, c)
}
