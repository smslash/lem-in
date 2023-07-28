package models

import "git/ssengerb/lem-in/internal/handle"

func NewGraph() *Graph {
	return &Graph{Rooms: make(map[string]*Node)}
}

func NewRoom(s []string) *Node {
	return &Node{
		Name:  s[0],
		X:     s[1],
		Y:     s[2],
		Both:  make(map[string]*Node),
		From:  make(map[string]*Node),
		To:    make(map[string]*Node),
		Level: 9223372036854775807,
	}
}

func (g *Graph) AddRoom(s []string) {
	r := NewRoom(s)
	g.Rooms[s[0]] = r

	if g.StartFound {
		g.Rooms[s[0]].StartRoom = true
		g.StartFound = false
	} else if g.EndFound {
		g.Rooms[s[0]].EndRoom = true
		g.EndFound = false
	}
}

func (g *Graph) AddEdge(room1, room2 string) {
	if room1 == room2 {
		return
	}

	firstRoom := g.GetRoom(room1)
	secondRoom := g.GetRoom(room2)

	firstRoom.Both[secondRoom.Name] = secondRoom
	secondRoom.Both[firstRoom.Name] = firstRoom
}

func (g *Graph) DeleteEdge(room1, room2 string) {
	firstRoom := g.GetRoom(room1)
	secondRoom := g.GetRoom(room2)

	delete(firstRoom.Both, secondRoom.Name)
	delete(secondRoom.Both, firstRoom.Name)
}

func (g *Graph) AddEdgeFrom(room1, room2 string) {
	firstRoom := g.GetRoom(room1)
	secondRoom := g.GetRoom(room2)

	firstRoom.From[secondRoom.Name] = secondRoom
	secondRoom.From[firstRoom.Name] = firstRoom
}

func (g *Graph) GetRoom(name string) *Node {
	var find bool
	for i, v := range g.Rooms {
		if v.Name == name {
			find = true
			return g.Rooms[i]
		}
	}
	if !find {
		handle.Error("ERROR: " + name + " room does not exist")
	}
	return nil
}

func (g *Graph) GetStartRoom() string {
	for _, v := range g.Rooms {
		if v.StartRoom {
			return v.Name
		}
	}
	return ""
}

func (g *Graph) GetEndRoom() string {
	for _, v := range g.Rooms {
		if v.EndRoom {
			return v.Name
		}
	}
	return ""
}
