package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(IsPalindrome("car"))                        
}

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	re := regexp.MustCompile(`[a-z0-9]+`)
	chars := strings.Join(re.FindAllString(s, -1), "")
	for i := 0; i < len(chars)/2; i++ {
		if chars[i] != chars[len(chars)-1-i] {
			return false
		}
	}
	return true
}


