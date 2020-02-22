package models

import (
	"royal-rumble/utils"
	"strings"
)

// RoyalNames holds list of royal names
type RoyalNames []string

// Len finds the number of names.
func (rn RoyalNames) Len() int {
	return len(rn)
}

// Less finds if element i is less then element j.
func (rn RoyalNames) Less(i, j int) bool {
	iSplit := strings.Split(rn[i], " ")
	iName := strings.Join(iSplit[:len(iSplit)-1], " ")
	iRromanNumber := iSplit[len(iSplit)-1]

	jSplit := strings.Split(rn[j], " ")
	jName := strings.Join(jSplit[:len(jSplit)-1], " ")
	jRromanNumber := jSplit[len(jSplit)-1]

	switch strings.Compare(iName, jName) {
	case -1:
		return true
	case 1:
		return false
	default:
		roman := utils.NewRomans()
		return roman.ToInt(iRromanNumber) < roman.ToInt(jRromanNumber)
	}
}

// Swap swaps elemtn i with element j
func (rn RoyalNames) Swap(i, j int) {
	rn[i], rn[j] = rn[j], rn[i]
}