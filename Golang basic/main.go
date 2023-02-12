package main

import "fmt"

//Create a function to reverse a string.
func reverse(s string) string{
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Create a function to check if a string is a palindrome.
func isPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// Create a function to calculate the factorial of a given number.
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Create a function to find the greatest common divisor of two numbers
func greatest_common_divisor(a int, b int) int{
	if b == 0 {
		return a
	}
	return greatest_common_divisor(b, a % b)
}

// Create a function to find the least common multiple of two numbers.
func least_common_multiple(a int, b int) int{
	return a * b / greatest_common_divisor(a, b)
}

// Create a function to check if a given number is a prime number.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

// Create a function to create a Fibonacci series up to a given number.
func fibonacci(n int) []int {
	fib := []int{0, 1}
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	return fib
}

// Create a function to bubble sort an array of numbers.
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// Create a function to calculate the sum of an array of numbers.
func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// Create a function to find the maximum element in an array of numbers
func max(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	// test reverse function
	fmt.Println(reverse("Hello, world"))

	// test isPalindrome function
	fmt.Println(isPalindrome("Hello, world"))
	fmt.Println(isPalindrome("HelloolleH"))

	// test factorial function
	fmt.Println(factorial(5))

	// test greatest_common_divisor function
	fmt.Println(greatest_common_divisor(10, 5))

	// test least_common_multiple function
	fmt.Println(least_common_multiple(10, 5))

	// test isPrime function
	fmt.Println(isPrime(2))
	fmt.Println(isPrime(10))

	// test fibonacci function
	fmt.Println(fibonacci(10))

	// test bubbleSort function
	fmt.Println(bubbleSort([]int{5, 4, 3, 2, 1}))

	// test sum function
	fmt.Println(sum([]int{5, 4, 3, 2, 1}))

	// test max function
	fmt.Println(max([]int{5, 4, 3, 2, 1}))
}