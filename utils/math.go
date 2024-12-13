package utils

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
