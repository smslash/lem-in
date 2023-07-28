package main

import (
	"os"

	"git/ssengerb/lem-in/internal/algo"
	"git/ssengerb/lem-in/internal/handle"
	"git/ssengerb/lem-in/internal/models"
	"git/ssengerb/lem-in/internal/parse"
)

func main() {
	if len(os.Args) != 2 {
		handle.Error("ERROR: invalid arguments\n\nUsage: go run . [FILE]\n\nEX: go run . example00.txt")
	}

	g := models.NewGraph()

	parse.OpenFile(os.Args[1], g)

	start := g.GetRoom(g.GetStartRoom())
	end := g.GetRoom(g.GetEndRoom())

	algo.BFS(start)

	if !end.Visited {
		handle.Error("ERROR: no way to go from \"" + start.Name + "\" - start to \"" + end.Name + "\" - end")
	}

	rotation := algo.Min(len(end.Both), len(end.Both))

	for i := 0; i < rotation; i++ {
		algo.NullVisited(g)
		algo.Bhandari(end)
	}

	algo.PrintAnswer(start, end, g)
}
