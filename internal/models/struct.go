package models

type Graph struct {
	Ants       int
	Rooms      map[string]*Node
	StartFound bool
	EndFound   bool
	NStart     bool
	NEnd       bool
}

type Node struct {
	Name      string
	X, Y      string
	Both      map[string]*Node
	From      map[string]*Node
	To        map[string]*Node
	Visited   bool
	Red       bool
	Level     int
	Busy      bool
	Ant       Ant
	StartRoom bool
	EndRoom   bool
}

type Ant struct {
	Number int
	Moved  bool
}
