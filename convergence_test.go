package ode

import (
	"math"
	"testing"
)

var (
	exp_hh = []float64{1e-1, 1e-2, 1e-3}
	exp_nn = []int{1e1, 1e2, 1e3}
)

func TestOrderOfConvergence(t *testing.T) {
	testCases := []struct {
		solver         Solver
		expected_order float64
	}{
		{
			solver:         ExplicitEuler,
			expected_order: 1,
		},
		{
			solver:         StandardRungeKutta(),
			expected_order: 4,
		},
	}

	for _, tc := range testCases {
		orders := OrderOfConvergence(
			tc.solver,
			exp_f,
			exp_t0,
			exp_x0,
			exp_x1,
			exp_hh,
			exp_nn,
		)
		for _, o := range orders {
			if math.Abs(tc.expected_order-o) > 0.1 {
				t.Errorf(
					"unexpected order of convergence: got %v, want [%v, ...]",
					o,
					tc.expected_order,
				)
			}
		}
	}
}
