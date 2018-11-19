package ode

import "gonum.org/v1/gonum/mat"

// ExplicitRungeKutta constructs an explicit Runge-Kutta method based on the
// Butcher tableau
//
//  c |  A
// ---+-----
//    | b^T
//
// Only the elements below the diagonal of A are accessed.
func ExplicitRungeKutta(
	A mat.Matrix,
	b, c mat.Vector,
) Solver {
	return func(
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

		r := b.Len()
		dim := x0.Len()
		k := make([]mat.Vector, r)

		for i := 0; i < n; i++ {
			k[0] = f(t, x[i])
			for j := 1; j < r; j++ {
				tt := t + c.AtVec(0)*h
				xx := mat.NewVecDense(dim, nil)
				for m := 0; m < j; m++ {
					xx.AddScaledVec(xx, A.At(j, m), k[m])
				}
				xx.AddScaledVec(x[i], h, xx)
				k[j] = f(tt, xx)
			}
			step := mat.NewVecDense(dim, nil)
			for j := 0; j < r; j++ {
				step.AddScaledVec(step, b.AtVec(j), k[j])
			}
			step.AddScaledVec(x[i], h, step)
			x[i+1] = step
			t += h
		}

		return x
	}
}
