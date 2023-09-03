package main

import (
	"fmt"
)

func Compare(a, b string) string {
	lenA := len(a)
	lenB := len(b)

	// find the longest common substring between string A and string B using a nested loop
	for i := 0; i < lenA; i++ {
		for j := 0; j < lenB; j++ {
			k := 0
			// if a character at index i+k in string A matches a character at index j+k in string B, increment k until a mismatch is found or until one of the strings runs out of character
			for i+k < lenA && j+k < lenB && a[i+k] == b[j+k] {
				k++
			}
			// if k > 0 and we've reached the end of one of the strings or found a character that doesn't match in string A or B, return the longest common substring
			if k > 0 && (i+k == lenA || j+k == lenB || a[i+k] != b[j+k]) {
				return a[i : i+k]
			}
		}
	}
	// if there is no common substring, return an empty string
	return ""
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
