package main

import "testing"

type Test struct {
	in  complex128
	out complex128
}

func run_tests(f Function, t *testing.T, tests []Test) {
	t.Log("testing", f)
	for _, test := range tests {
		val := f.Evaluate(test.in)
		t.Logf("testing at %v, got %v", test.in, val)
		if val != test.out {
			t.Errorf("did not evaluate to expected %v at %v", test.out, test.in)
		}
	}
}

func TestConstantZero(t *testing.T) {
	c := constant(0)
	tests := []Test{
		{0, 0},
		{1, 0},
		{1i, 0},
		{1 + 1i, 0},
		{200 - 50i, 0},
	}
	run_tests(c, t, tests)
}

func TestPolyIntegers(t *testing.T) {
	f := NewPoly(1, 2, 1)
	tests := []Test{
		{0, 1},
		{1, 4},
		{1i, 2i},
		{1 + 1i, 3 + 4i},
		{2 - 3i, -18i},
	}
	run_tests(f, t, tests)
}
