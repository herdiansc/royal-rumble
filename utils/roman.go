package utils

import "strings"

// Romans holds list of mapping string roman number to numerical number.
type Romans map[string]int

// NewRomans initiates Romans.
func NewRomans() Romans {
	return Romans{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
	}
}

// ToInt converts roman number to numerical number based on Romans data.
func (r Romans) ToInt(letter string) int {
	n := 0
	var sign int
	chars := strings.Split(letter, "")
	for i := range chars {
		if _, ok := r[chars[i]]; !ok {
			return 0
		}
		sign = 1
		if i+1 < len(chars) && r[chars[i+1]] > r[chars[i]] {
			sign = -1
		}
		n = n + (r[chars[i]] * sign)
	}
	return n
}
