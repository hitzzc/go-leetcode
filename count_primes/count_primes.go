package count_primes

func countPrimes(n int) int {
	check := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if check[i] {
			continue
		}
		for j := i; i*j < n; j++ {
			if !check[i*j] {
				check[i*j] = true
			}
		}
	}
	cnt := 0
	for i := 2; i < n; i++ {
		if !check[i] {
			cnt++
		}
	}
	return cnt
}
