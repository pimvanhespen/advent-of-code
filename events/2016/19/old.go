package main

// This file contains older versions of the solution, which are kept for reference.

// initial naive solution - works in O(2*N)
func solveNaive(n int) int {
	if n <= 1 {
		return n
	}

	var nums []int

	if n%2 == 0 {
		nums = make([]int, 0, n/2)
		for i := 1; i <= n; i += 2 {
			nums = append(nums, i)
		}
	} else {
		nums = make([]int, 0, n/2)
		for i := 3; i <= n; i += 2 {
			nums = append(nums, i)
		}
	}

	for len(nums) > 1 {
		offset := len(nums) % 2

		for i := len(nums) - 1 - offset; i >= 0; i -= 2 {
			nums = append(nums[:i], nums[i+1:]...)
		}

		if offset == 1 {
			nums = nums[1:]
		}
	}

	return nums[0]
}

// this solution works in O(1+)
func solveQuick(n int) int {
	in := uint32(n)
	for i := 31; i >= 0; i-- {
		mask := uint32(1 << uint(i))
		if in&mask != 0 {
			return int(1 + 2*(in-mask))
		}
	}

	return 0
}
