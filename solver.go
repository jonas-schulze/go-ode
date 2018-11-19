package ode

import "gonum.org/v1/gonum/mat"

type Solver func(
	// derivative of x
	f Rhs,
	// initial condition: x(t0) == x0
	t0 float64, x0 mat.Vector,
	// time delta for equidistant interpolation points
	h float64,
	// number of time steps to stop after
	n int,
) []mat.Vector
