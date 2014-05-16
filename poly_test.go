package main

import "testing"

type EvalData struct {
	in  complex128
	out complex128
}

type DegreeData struct {
	in  poly
	out int
}

type NormData struct {
	in  []complex128
	out []complex128
}

func run_tests(f Function, t *testing.T, tests []EvalData) {
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
	tests := []EvalData{
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
	tests := []EvalData{
		{0, 1},
		{1, 4},
		{1i, 2i},
		{1 + 1i, 3 + 4i},
		{2 - 3i, -18i},
	}
	run_tests(f, t, tests)
}

func TestDegree(t *testing.T) {
	tests := []DegreeData{
		{poly([]complex128{}), 0},
		{poly([]complex128{0}), 0},
		{poly([]complex128{1}), 0},
		{poly([]complex128{0, 1}), 1},
		{poly([]complex128{1, 1, 0}), 1},
	}
	for _, test := range tests {
		deg := test.in.Degree()
		t.Logf("deg %v is %v", test.in, deg)
		if deg != test.out {
			t.Errorf("did not get expected degree %v for %v", test.out, test.in)
		}
	}
}

func TestNormalise(t *testing.T) {
	tests := []NormData{
		{[]complex128{}, []complex128{}},
		{[]complex128{0}, []complex128{}},
		{[]complex128{0, 0}, []complex128{}},
		{[]complex128{1}, []complex128{1}},
		{[]complex128{1, 0, 1}, []complex128{1, 0, 1}},
		{[]complex128{1, 0, 1, 0}, []complex128{1, 0, 1}},
		{[]complex128{1, 1, 0}, []complex128{1, 1}},
	}
	for _, test := range tests {
		in := poly(test.in)
		out := poly(test.out)
		norm := in.Normalise()
		t.Logf("(%v).Normalise() is %v", in, norm)
		if len(out) != len(norm) {
			t.Error("did not get expected", out)
		}
		for i := range out {
			if out[i] != norm[i] {
				t.Error("did not get expected", out)
			}
		}
	}
}
