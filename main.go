package main

import (
	"GolangAStar/astar"
	cw "TCellConsoleWrapper"
	"strconv"
)

func main() {
	cw.Init_console()
	defer cw.Close_console()
	mymap := []string{
		"@..............",
		"..###..........",
		"...............",
		"...............",
		"........#......",
		"........#......",
		"........#......",
		"...............",
		"...............",
		"...............",
		"....###########",
		"....#....#.....",
		" ####.##.#.#.#.",
		"..#...#..#.#.#.",
		"....#.#....#..X",
	}
	costmap := getCostMapFromStringList(&mymap)
	key := ""
	fromx, fromy := 1, 1
	tox, toy := len(mymap) - 1, len(mymap[0]) - 1
	for key != "ESCAPE" {
		path := astar.FindPath(costmap, fromx, fromy, tox, toy)
		cw.SetFgColor(cw.BEIGE)
		for x := 0; x < len(mymap); x++ {
			for y := 0; y < len(mymap[0]); y++ {
				cw.PutChar(rune(mymap[x][y]), x, y)
			}
		}
		cw.SetFgColor(cw.MAGENTA)
		for _, c := range *path {
			cw.PutChar('*', c.X, c.Y)
		}
		cw.SetFgColor(cw.GREEN)
		cw.PutChar('@', fromx, fromy)
		cw.SetFgColor(cw.RED)
		cw.PutChar('X', tox, toy)
		cw.Flush_console()
		key = cw.ReadKey()
		switch key {
		case "2":
			fromy += 1
		case "4":
			fromx -= 1
		case "8":
			fromy -= 1
		case "6":
			fromx += 1

		}
	}
}

func getCostMapFromStringList(strmap *[]string) *[][]int {
	width, height := len(*strmap), len((*strmap)[0])
	costmap := make([][]int, height)
	for i := range costmap {
		costmap[i] = make([]int, width)
	}
	for i:=0; i<width; i++ {
		for j:=0; j<height; j++ {
			if (*strmap)[i][j] == '#' {
				costmap[i][j] = -1
			}
		}
	}
	cw.PutString(strconv.Itoa(len(costmap)) + " " + strconv.Itoa(len(costmap[0])), 0, 20)
	cw.Flush_console()
	return &costmap
}

