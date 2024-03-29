package recursion

func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * fact(n-1)
}
