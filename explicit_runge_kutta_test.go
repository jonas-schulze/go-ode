package ode

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

var (
	exp_f  = func(_ float64, x mat.Vector) mat.Vector { return x }
	exp_x0 = mat.NewVecDense(1, []float64{1})
)

const (
	exp_t0 = 0.0
	exp_h  = 1e-4
	exp_n  = 1e4
)

func BenchmarkExplicitEuler(b *testing.B) {
	s := ExplicitEuler
	for i := 0; i < b.N; i++ {
		s(exp_f, exp_t0, exp_x0, exp_h, exp_n)
	}
}
func BenchmarkExplicitEulerAsRungeKutta(b *testing.B) {
	aa := mat.NewDense(1, 1, []float64{0})
	bb := mat.NewVecDense(1, []float64{1})
	cc := mat.NewVecDense(1, []float64{0})
	s := ExplicitRungeKutta(aa, bb, cc)
	for i := 0; i < b.N; i++ {
		s(exp_f, exp_t0, exp_x0, exp_h, exp_n)
	}
}
