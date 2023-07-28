package parse

import (
	"bufio"
	"os"
	"strings"

	"git/ssengerb/lem-in/internal/handle"
	"git/ssengerb/lem-in/internal/models"
)

func parseData(file *os.File, g *models.Graph) {
	scanner := bufio.NewScanner(file)
	var queue int

	for scanner.Scan() {
		str := scanner.Text()
		if isEmpty(str) || isComment(str) {
			continue
		}

		line := strings.Fields(str)
		if queue == 0 && len(line) == 1 {
			ants(line[0], g)
			queue++
		} else if queue == 1 {
			queue = rooms(line, g)
		} else if queue == 2 && len(line) == 1 && isLink(line[0]) {
			g.AddEdge(findRoomsName(line[0]))
		} else {
			handle.Error("ERROR: invalid data format")
		}
	}

	if !existStart(g) {
		handle.Error("ERROR: start room not found")
	} else if !existEnd(g) {
		handle.Error("ERROR: end room not found")
	}

	if !isValidCoord(g) {
		handle.Error("ERROR: invalid coordinates")
	}

	if err := scanner.Err(); err != nil {
		handle.Error("ERROR: during scanning")
	}
}

func existStart(g *models.Graph) bool {
	for _, v := range g.Rooms {
		if v.StartRoom {
			g.StartFound = true
			return true
		}
	}
	return false
}

func existEnd(g *models.Graph) bool {
	for _, v := range g.Rooms {
		if v.EndRoom {
			g.EndFound = true
			return true
		}
	}
	return false
}
