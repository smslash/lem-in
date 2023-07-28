package algo

import (
	"git/ssengerb/lem-in/internal/models"
)

func FinalPaths(path []*models.Node, node *models.Node) []*models.Node {
	path = append(path, node)
	if node.EndRoom {
		return path
	} else {
		for _, v := range node.To {
			return FinalPaths(path, v)
		}
	}
	return []*models.Node{}
}

func NullVisited(g *models.Graph) {
	for _, v := range g.Rooms {
		v.Visited = false
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
