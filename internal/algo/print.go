package algo

import (
	"fmt"
	"strconv"

	"git/ssengerb/lem-in/internal/models"
)

func PrintGraph(g *models.Graph) {
	for i, v := range g.Rooms {
		fmt.Println("Name:", i)
		for j := range v.To {
			if len(v.To) != 0 {
				fmt.Print(j + " ")
			}
		}
		fmt.Println()
	}
}

func PrintPath(path []*models.Node) {
	if len(path) == 0 {
		return
	}
	for i := 0; i < len(path); i++ {
		fmt.Print(path[i].Name + "  ")
		if !path[i].EndRoom {
			fmt.Print("->  ")
		}
	}
	fmt.Println()
}

func PrintLevel(g *models.Graph) {
	for i, v := range g.Rooms {
		fmt.Printf("Name: %v \t Level: %v\n", i, v.Level)
	}
}

func PrintAnswer(start, end *models.Node, g *models.Graph) {
	answer := make([][]*models.Node, len(start.To))
	res := make([]*models.Node, 0)
	index := 0

	for _, v := range start.To {
		res = append(answer[index], start)
		answer[index] = FinalPaths(res, v)
		index++
	}

	SortPaths(answer)

	// for i := 0; i < len(answer); i++ {
	// 	PrintPath(answer[i])
	// }

	// fmt.Print("\n\n")

	number := 0
	str2 := ""

	for i := 0; i < len(answer); i++ {
		if len(answer[i]) == 2 {
			for number != g.Ants {
				fmt.Printf("L%v-%v ", number+1, end.Name)
				number++
			}
			fmt.Println()
			return
		}
	}

	for {
		str := str2
		if len(str2) > 0 {
			str2 = ""
		}
		written := false

		for i := 0; i < len(answer); i++ {
			if number == g.Ants {
				break
			}
			answer[i][1].Busy = true
			answer[i][1].Ant.Number = number + 1
			number++
		}

		for i := 0; i < len(answer); i++ {
			if len(answer[i]) == 2 {
				n := strconv.Itoa(answer[i][1].Ant.Number)
				str += "L" + n + "-" + answer[i][1].Name + " "
				continue
			}
			for j := len(answer[i]) - 2; j > 0; j-- {
				if answer[i][j].Busy && answer[i][j].Ant.Number != 0 && !answer[i][j].Ant.Moved {
					n := strconv.Itoa(answer[i][j].Ant.Number)
					str += "L" + n + "-" + answer[i][j].Name + " "

					written = true
					for _, v := range answer[i][j].To {
						v.Ant.Number = answer[i][j].Ant.Number
						v.Busy = true
						v.Ant.Moved = true
					}

					answer[i][j].Busy = false

					for _, p := range answer[i][j].To {
						if p == end {
							str2 += "L" + n + "-" + p.Name + " "
						}
					}
				}
			}
		}

		fmt.Print(SortString(str))

		for i := 0; i < len(answer); i++ {
			for j := 0; j < len(answer[i]); j++ {
				answer[i][j].Ant.Moved = false
			}
		}

		if !written {
			break
		}

		fmt.Println()
	}
	fmt.Println()
}
