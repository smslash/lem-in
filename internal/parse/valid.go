package parse

import (
	"strings"

	"git/ssengerb/lem-in/internal/handle"
	"git/ssengerb/lem-in/internal/models"
)

func isComment(s string) bool {
	if s[0] == '#' && s != "##start" && s != "##end" {
		return true
	}
	return false
}

func isEmpty(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\r' {
			return false
		}
	}
	return true
}

func isLink(s string) bool {
	return strings.Contains(s, "-")
}

func findRoomsName(s string) (string, string) {
	for i := 0; i < len(s); i++ {
		if s[i] == '-' && i+1 < len(s) {
			if isValidName(s[:i], s[i+1:]) {
				return s[:i], s[i+1:]
			}
			handle.Error("ERROR: room name can not start with \"L\" or have spaces")
		}
	}
	return "", ""
}

func isValidCoord(g *models.Graph) bool {
	for _, v := range g.Rooms {
		for _, j := range v.Both {
			if v.X == j.X && j.Y == v.Y {
				return false
			}
		}
	}
	return true
}

func isValidName(s1, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}

	if s1[0] == 'L' || s2[0] == 'L' {
		return false
	}

	space := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == ' ' {
			space++
		}
	}
	for i := 0; i < len(s2); i++ {
		if s2[i] == ' ' {
			space++
		}
	}

	if space > 0 {
		return false
	}

	return true
}
