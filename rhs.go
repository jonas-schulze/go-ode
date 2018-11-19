package ode

import "gonum.org/v1/gonum/mat"

// Rhs defines a right-hand side of an ODE, that is the derivative of x.
type Rhs func(t float64, x mat.Vector) mat.Vector
