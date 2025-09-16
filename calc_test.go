package main

import "testing"

func TestAdd(t *testing.T) {
	if Add(2, 3) != 5 {
		t.Errorf("Add(2,3) should be 5")
	}
}

func TestSub(t *testing.T) {
	if Sub(5, 3) != 2 {
		t.Errorf("Sub(5,3) should be 2")
	}
}

func TestMul(t *testing.T) {
	if Mul(2, 3) != 6 {
		t.Errorf("Mul(2,3) should be 6")
	}
}

func TestDiv(t *testing.T) {
	if Div(6, 3) != 2 {
		t.Errorf("Div(6,3) should be 2")
	}
	if Div(6, 0) != 0 {
		t.Errorf("Div(6,0) should be 0 (avoid panic)")
	}
}
