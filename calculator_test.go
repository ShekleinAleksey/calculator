package main

import "testing"

func TestAdd(t *testing.T) {
	var exp float64
	exp = 5
	res := addition(2, 3)
	if res != exp {
		t.Errorf("%f was expect, but got %f \n", exp, res)
	}
}

func TestSubstract(t *testing.T) {
	var exp float64
	exp = 2
	res := substract(8, 6)
	if res != exp {
		t.Errorf("%f was expectbut got %f \n", exp, res)
	}
}

func TestMultiply(t *testing.T) {
	var exp float64
	exp = 10
	res := multiply(2, 5)
	if res != exp {
		t.Errorf("%f was expectbut got %f \n", exp, res)
	}
}

func TestDivision(t *testing.T) {
	var exp float64
	exp = 2
	res := division(6, 3)
	if res != exp {
		t.Errorf("%f was expectbut got %f \n", exp, res)
	}
}
