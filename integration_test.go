package numericIntegration

import "testing"
import "math"

func TestIntegrateConstFunc(t *testing.T) {
	var f = func(x float64) float64 {
		return 1
	}
	val := Integrate(f, 0, 1)
	if !arefloatsEqual(val, 1) {
		t.Error("expected 1, got ", val)
	}

}

func TestIntegrateSinFunc(t *testing.T) {
	var f = func(x float64) float64 {
		return math.Sin(x)
	}
	val := Integrate(f, 0, math.Pi)
	if !arefloatsEqual(val, 2) {
		t.Error("expected 2, got ", val)
	}
}

func arefloatsEqual(f1 float64, f2 float64) bool {
	eps := 0.00001
	return f1 >= f2-eps && f1 <= f2+eps
}
