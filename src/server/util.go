package server

import "strings"

// checkIfPalindrome checks to see if the given string is a palindrome
func checkIfPalindrome(s string) bool {
	// an empty string or a string with a single character is always a palindrome
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	// two indices to keep track of each character from left and right
	var left, right = 0, len(s) - 1
	for left <= right {
		if strings.ToLower(string(s[right])) < string('a') ||
			strings.ToLower(string(s[right])) > string('z') { // jump over the non-alphabet character from right
			right--
		} else if strings.ToLower(string(s[left])) < string('a') ||
			strings.ToLower(string(s[left])) > string('z') { // jump over the non-alphabet character from left
			left++
		} else if strings.ToLower(string(s[right])) !=
			strings.ToLower(string(s[left])) { // characters at both ends don't match
			return false
		} else { // characters at both ends match
			left++
			right--
		}
	}
	return true
}
