package main

type Algorithm struct {}

/* print sum of two numbers */
func (a *Algorithm) sum(x int, y int) int {
	return x+y
}

/* print multiplication of two numbers */
func (a *Algorithm) multiply(x int, y int) int {
	return x*y
}

/* print first N prime numbers */
func (a *Algorithm) prime(N int) []int {
	count := 0
	num := 2
	var prime []int
	for count < N {
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			prime = append(prime, num)
			count++
		}
		num++
	}
	return prime
}

/* print first N fibonacci numbers */
func (a *Algorithm) fib(N int) []int {
	fib := make([]int, N)

	fib[0] = 0
	fib[1] = 1

	for i:=2; i<N; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[:N]
}