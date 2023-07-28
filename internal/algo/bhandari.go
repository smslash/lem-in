package algo

import (
	"git/ssengerb/lem-in/internal/models"
)

var MAX = 9223372036854775807

func Bhandari(end *models.Node) {
	path := make([]*models.Node, 0)
	path = append(path, end)
	end.Red = true
	end.Visited = true

	min := MAX
	for _, v := range end.Both {
		if !v.Visited && !v.Red && v.Level < min {
			min = v.Level
		}
	}

	for _, v := range end.Both {
		if !v.Visited && !v.Red && v.Level == min {
			v.Red = true
			v.Visited = true
			path = append(path, v)
			v.To[end.Name] = end
			path, _ = findPath(path, v)
			break
		}
	}
}

func findPath(path []*models.Node, node *models.Node) ([]*models.Node, bool) {
	if node.EndRoom {
		return path[:len(path)-1], false
	}

	if node.StartRoom {
		if !path[len(path)-1].StartRoom {
			path = append(path, node)
			node.To[path[len(path)-2].Name] = path[len(path)-2]
		}
		return path, true
	}

	for _, v := range node.Both {
		var f bool
		min := MAX
		red_neighbors := 0
		for _, v := range node.Both {
			if !v.Visited && !v.Red && v.Level <= min {
				min = v.Level
			}
			if v.Red {
				red_neighbors++
			}
		}

		if red_neighbors == len(node.Both) {
			path, f = redPath(path, node)
			return path, f
		}

		if v.StartRoom {
			path = append(path, v)
			v.To[node.Name] = node
			return path, true
		}

		if !v.Visited && !v.Red && v.Level == min {
			v.Red = true
			v.Visited = true
			path = append(path, v)
			v.To[node.Name] = node
			path, f = findPath(path, v)
			if !f {
				v.Red = false
				path = path[:len(path)-1]
				delete(v.To, node.Name)
				continue
			} else {
				return path, true
			}
		}
	}

	return path, false
}

func redPath(path []*models.Node, node *models.Node) ([]*models.Node, bool) {
	for _, v := range node.Both {
		var f bool
		min := MAX
		var tmp *models.Node
		var save *models.Node
		for _, v := range node.Both {
			if v.StartRoom {
				path = append(path, v)
				v.To[node.Name] = node
				return path, true
			}

			if !v.Visited && v.Level <= min {
				min = v.Level
			}
		}

		if !v.Visited && v.Level == min {
			v.Visited = true
			path = append(path, v)
			if len(v.To) >= 1 {
				for _, j := range v.To {
					delete(v.To, j.Name)
					tmp = j
					save = j
				}
			}
			v.To[node.Name] = node
			path, f = findPath(path, tmp)
			if !f {
				for _, j := range v.To {
					delete(v.To, j.Name)
				}
				v.To[save.Name] = save
				continue
			} else {
				return path, true
			}
		}
	}

	return path, false
}
