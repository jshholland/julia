package main

import "testing"

func check_eval(f Function, t *testing.T, points, vals []complex128) {
	if len(points) != len(vals) {
		panic("len(points) != len(vals)")
	}

	for i, z := range points {
		val := f.Evaluate(z)
		t.Logf("testing at %v, got %v", z, val)
		if val != vals[i] {
			t.Errorf("did not evaluate to expected %v at %v", vals[i], z)
		}
	}
}

func TestConstantZero(t *testing.T) {
	c := constant(0)
	points := []complex128{0, 1, 1i, 1 + 1i, 200 - 50i}
	vals := []complex128{0, 0, 0, 0, 0}
	check_eval(c, t, points, vals)
}

func TestPolyIntegers(t *testing.T) {
	f := NewPoly(1, 2, 1)
	points := []complex128{0, 1, 1i, 1 + 1i, 2 - 3i}
	vals := []complex128{1, 4, 2i, 3 + 4i, -18i}
	check_eval(f, t, points, vals)
}
