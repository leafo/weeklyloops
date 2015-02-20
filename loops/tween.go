package loops

func Smoothstep(t float64) float64 {
	return t * t * t * (t*(t*6.0-15.0) + 10.0)
}
