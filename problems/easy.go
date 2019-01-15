package problems

import (
	"math"
)

// Contains solutions for easy problems: https://leetcode.com/problemset/all/?difficulty=Easy

// https://leetcode.com/problems/hamming-distance/
// hammingDistance calculates the Hamming distance.
// The Hamming distance between two integers is the number of positions
// at which the corresponding bits are different.
func hammingDistance(x int, y int) int {
	var distance int
	// XOR of two numbers
	// 6(110) and 0(000) gives 6(110)
	value := x ^ y
	// count the number of bits set
	for value != 0 {
		// A bit is set, so increment the count and clear the bit
		distance++
		// counting bits by algorihm of Peter Wegner
		// http://dl.acm.org/citation.cfm?id=367236.367286
		// step1: 6(110) & 5(101) = 4(100) (distance=1)
		// step2: 4(100) & 3(011) = 0(000) (distance=2)
		value &= value - 1
	}
	// Return the number of differing bits
	return distance
}

// https://leetcode.com/problems/single-number/
// SingleNumber finds single integer.
// Given an array of integers, every element appears twice except for one.
// Note:Your algorithm should have a linear runtime complexity.
// Could you implement it without using extra memory?
func singleNumber(numbers []int) int {
	var result int
	for _, number := range numbers {
		result = result ^ number
	}

	return result
}

// https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	cache := make(map[int]int, len(nums))
	for i, num := range nums {
		complement := target - num
		if firstIndex, ok := cache[complement]; ok {
			return []int{firstIndex, i}
		}

		cache[num] = i
	}

	return nil
}

// https://leetcode.com/problems/reverse-integer/
func reverse(x int) int {
	x64 := int64(x)
	sign := 1
	if x64 < 0 {
		sign = -1
		x64 = -x64
	}

	var reverted int64
	for ; x64 != 0; x64 /= 10 {
		reverted = reverted*10 + x64%10
		if reverted > math.MaxInt32 {
			return 0
		}
	}

	return int(reverted) * sign
}

// https://leetcode.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x != 0 && x%10 == 0 {
		return false
	}

	var reverted int
	for ; x > reverted; x /= 10 {
		reverted = reverted*10 + x%10
	}

	return x == reverted || x == reverted/10
}

// https://leetcode.com/problems/contains-duplicate-ii/
func containsNearbyDuplicate(nums []int, k int) bool {
	cache := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := cache[num]; ok {
			if i-j <= k {
				return true
			}
		}

		cache[num] = i
	}

	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pValue := p.Val
	qValue := q.Val

	for node := root; node != nil; {
		parentValue := node.Val

		switch {
		case pValue > parentValue && qValue > parentValue:
			node = node.Right
		case pValue < parentValue && qValue < parentValue:
			node = node.Left
		default:
			return node
		}
	}

	return nil
}

// https://leetcode.com/problems/self-dividing-numbers/
func selfDividingNumbers(left int, right int) []int {
	var selfDivNumbers []int
	for i := left; i <= right; i++ {
		if ok := isSelfDiv(i); ok {
			selfDivNumbers = append(selfDivNumbers, i)
		}
	}
	return selfDivNumbers
}

func isSelfDiv(num int) bool {
	for x := num; x > 0; x /= 10 {
		lastDigit := x % 10
		if lastDigit == 0 || (num%lastDigit) > 0 {
			return false
		}
	}

	return true
}

// https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	result := getRomanValue(s[length-1])

	for i := length - 2; i >= 0; i-- {
		value := getRomanValue(s[i])

		if value < getRomanValue(s[i+1]) {
			result -= value
			continue
		}

		result += value
	}

	return result
}

func getRomanValue(b byte) int {
	switch b {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	var stack []rune
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}

			lastIndex := len(stack) - 1
			top := stack[lastIndex]
			stack = stack[:lastIndex]

			if r == ')' && top != '(' ||
				r == ']' && top != '[' ||
				r == '}' && top != '{' {
				return false
			}
		}
	}

	return len(stack) == 0
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var root ListNode
	tail := &root

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}

		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	return root.Next
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] == nums[i] {
			continue
		}

		i++
		nums[i] = nums[j]
	}

	return i + 1
}
