package algo

import (
	"strconv"
	"strings"

	"git/ssengerb/lem-in/internal/handle"
	"git/ssengerb/lem-in/internal/models"
)

func SortPaths(s [][]*models.Node) [][]*models.Node {
	for i := 0; i < len(s)-1; i++ {
		swapped := false
		for j := 0; j < len(s)-i-1; j++ {
			if len(s[j+1]) < len(s[j]) {
				s[j], s[j+1] = s[j+1], s[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return s
}

func SortString(s string) string {
	str := strings.Fields(s)
	res := ""

	for i := 0; i < len(str)-1; i++ {
		swapped := false
		for j := 0; j < len(str)-i-1; j++ {
			if compareNumber(str[j+1], str[j]) {
				str[j], str[j+1] = str[j+1], str[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	for i := 0; i < len(str); i++ {
		res += str[i]
		if i+1 < len(str) {
			res += " "
		}
	}

	return res
}

func compareNumber(s1, s2 string) bool {
	tmp1 := ""
	for i := 1; s1[i] != '-'; i++ {
		tmp1 += string(s1[i])
	}
	n1, err := strconv.Atoi(tmp1)
	if err != nil {
		handle.Error("ERROR: can not convert " + tmp1 + " to string")
	}

	tmp2 := ""
	for i := 1; s2[i] != '-'; i++ {
		tmp2 += string(s2[i])
	}
	n2, err := strconv.Atoi(tmp2)
	if err != nil {
		handle.Error("ERROR: can not convert " + tmp2 + " to string")
	}

	if n2 < n1 {
		return false
	}

	return true
}
