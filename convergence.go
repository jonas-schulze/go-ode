package ode

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func OrderOfConvergence(
	s Solver,
	f Rhs,
	t0 float64,
	// x1 is the expected/exact end value, x(t0+n*h)
	x0, x1 mat.Vector,
	h []float64,
	n []int,
) []float64 {
	orders := make([]float64, 0, len(h)-1)
	old_logErr := 0.0
	old_logH := 0.0
	for i := range h {
		current_n := n[i]
		current_h := h[i]
		x_h := s(f, t0, x0, current_h, current_n)
		var err mat.VecDense
		err.SubVec(x_h[current_n], x1)
		logErr := math.Log(mat.Norm(&err, 2))
		logH := math.Log(current_h)

		if i > 0 {
			c := logErr - old_logErr
			c /= (logH - old_logH)
			orders = append(orders, c)
		}
		old_logErr = logErr
		old_logH = logH
	}

	return orders
}
