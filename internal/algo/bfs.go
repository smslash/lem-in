package algo

import (
	"container/list"

	"git/ssengerb/lem-in/internal/models"
)

func BFS(start *models.Node) {
	queue := list.New()

	queue.PushBack(start)
	start.Visited = true
	start.Level = 0

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		room := element.Value.(*models.Node)
		for _, neighbor := range room.Both {
			if !neighbor.Visited {
				queue.PushBack(neighbor)
				neighbor.Level = room.Level + 1
				neighbor.Visited = true
			}
		}
	}
}
