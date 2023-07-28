package parse

import (
	"strconv"

	"git/ssengerb/lem-in/internal/handle"
	"git/ssengerb/lem-in/internal/models"
)

func ants(s string, g *models.Graph) {
	var err error
	g.Ants, err = strconv.Atoi(s)
	if err != nil {
		handle.Error("ERROR: can not convert \"" + s + "\"")
	} else if g.Ants <= 0 {
		handle.Error("ERROR: number of ants can not be 0 or less")
	}
}

func rooms(s []string, g *models.Graph) int {
	if len(s) == 1 {
		if s[0] == "##start" {
			if g.NStart {
				handle.Error("ERROR: too many start rooms found")
			}
			g.NStart, g.StartFound = true, true
		} else if s[0] == "##end" {
			if g.NEnd {
				handle.Error("ERROR: too many end rooms found")
			}
			g.NEnd, g.EndFound = true, true
		} else if isComment(s[0]) {
			return 1
		} else if isLink(s[0]) {
			g.AddEdge(findRoomsName(s[0]))
			return 2
		} else {
			handle.Error("ERROR: invalid data format")
		}
	} else if len(s) == 3 {
		g.AddRoom(s)
	} else {
		handle.Error("ERROR: invalid data format")
	}
	return 1
}
