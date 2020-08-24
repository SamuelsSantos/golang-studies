package main

// Fatorial method to calc a fatorial number
func Fatorial(n int) int {

	result := 1
	for i := 1; i <= n; i++ {

		result *= i

	}
	return result;

}