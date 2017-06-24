package numericIntegration

func Integrate(f func(float64) float64, a float64, b float64) float64 {
	amtSteps := 1000.0
	len := b - a
	stepSize := len / amtSteps
	res := 0.0
	for leftSide := a; leftSide < b; leftSide += stepSize {
		rightSide := leftSide + stepSize
		res += (rightSide - leftSide) / 6 * (f(leftSide) + 4*f((leftSide+rightSide)/2) + f(rightSide))
	}
	return res
}
