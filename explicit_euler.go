package ode

import "gonum.org/v1/gonum/mat"

type Rhs func(t float64, x mat.Vector) mat.Vector

func ExplicitEuler(
	// derivative of x
	f Rhs,
	// initial condition: x(t0) == x0
	t0 float64, x0 mat.Vector,
	// time delta for equidistant interpolation points
	h float64,
	// number of time steps to stop after
	n int,
) []mat.Vector {
	// TODO: create raw data for []mat.Vector contiguously.
	x := make([]mat.Vector, n+1, n+1)
	x[0] = x0
	t := t0
	for i := 0; i < n; i++ {
		// x_{i+1} = x_i + h*f(t_i,x_i)
		step := f(t, x[i])
		var next mat.VecDense
		next.AddScaledVec(x[i], h, step)
		x[i+1] = &next
		t += h
	}
	return x
}
