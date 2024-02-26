func sum_to_n_a(n int) int {
  	if n < 1 {
    	return 0
  	}
  	if n == 1 {
		return 1
  	}
	return n + sum_to_n_a(n - 1)
	// This is a recursive method to sum to n
	// The time complexity is O(n)
	// If the compiler does no tail-optimization to this method, then the space complexity is also O(n) due to the stack of recursive calls
	// Otherwise, the space complexity is O(1)
}

func sum_to_n_b(n int) int {
  	sum := 0
	for i:=1;i<=n;i++ {
 		sum += i
  	}
  	return sum
	// Time complexity: O(n)
	// Space complexity: O(1)
}

func sum_to_n_c(n int) int {
	if n < 1 {
		return 0
	}
	m := n + 1
	var result int
	if m % 2 == 0 {
		result = (m / 2) * n
	} else {
		result = (n / 2) * m
	}
	return result
	// Time complexity: O(1)
	// Space complexity: O(1)
	// We don't just return (n + 1) * n / 2, just in case of overflow
}
