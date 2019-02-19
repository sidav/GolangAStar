package main

import (
	cw "TCellConsoleWrapper"
	"astar/astar"
	"strconv"
	"time"
)

func main() {
	cw.Init_console()
	defer cw.Close_console()
	mymap := []string{
		" ####################",
		"    #       #       #",
		"# # ####### # #######",
		"# # #     # #       #",
		"# # # ### # ### ### #",
		"# #   #   #   # #   #",
		"####### ##### ### # #",
		"#       #         # #",
		"# ### ### ### ##### #",
		"#   # #   # # # #   #",
		"### # # # # # # # ###",
		"#   #   #   #   # # #",
		"# ############# # # #",
		"#     #   #     #   #",
		"### ### # # ##### ###",
		"# # #   #   #   # # #",
		"# # # ######### # # #",
		"#   #     #     # # #",
		"# ####### # ##### # #",
		"#   #   #       #   #",
		"### # # ####### # ###",
		"#   # # #       #   #",
		"# ### # #############",
		"# # # #   #     #   #",
		"# # # ### ##### # # #",
		"# #   # #   # #   # #",
		"# # ### ### # # ### #",
		"# #   #     # # # # #",
		"# ### # # ### # # # #",
		"#     # #       #   #",
		"####### ########### #",
		"#   #     #   #   # #",
		"# # ##### # # ### ###",
		"# #       # #   #   #",
		"# ############# ### #",
		"#   #         #     #",
		"### # # ##### ##### #",
		"# #   #   #       # #",
		"# ### ### # ####### #",
		"#     #   #         #",
		"####### #############",
		"#       #       #   #",
		"# ### # # ##### # ###",
		"#   # # # #     #   #",
		"### # # # # ####### #",
		"# # # # # # #   #   #",
		"# # # # # # # # # ###",
		"#   # #   #   # # # #",
		"##### ######### # # #",
		"# #   #       #   # #",
		"# # ### # ######### #",
		"# #     #     #     #",
		"# ########### # #####",
		"#             # #   #",
		"##### # ### ### ### #",
		"#   # #   #         #",
		"# # ##### ### # #####",
		"# #   # # #   #     #",
		"##### # # # ### #####",
		"#     #   # #   #   #",
		"# ######### # ### # #",
		"#           # #   # #",
		"### # ####### # ### #",
		"#   # #   # # # #   #",
		"# ##### # # # # ### #",
		"#       # # #     # #",
		"# ####### # ##### ###",
		"# #         #   #   #",
		"# ### ####### # ### #",
		"#   # # #   # # #   #",
		"### # # # ### ### # #",
		"#   #   #   #     # #",
		"# ######### ####### #",
		"#   #   # #       # #",
		"### ### # ####### # #",
		"# #     # #         #",
		"# ### ### # #########",
		"#     #         #   #",
		"# # ####### ### # # #",
		"# #       #   # # # #",
		"###################  ",
	}
	costmap := getCostMapFromStringList(&mymap)
	key := ""
	fromx, fromy := 1, 1
	tox, toy := len(mymap) - 1, len(mymap[0]) - 1
	for key != "ESCAPE" {
		cw.SetFgColor(cw.DARK_GRAY)
		startTime := time.Now()
		path := astar.FindPath(costmap, fromx, fromy, tox, toy, true, true)
		cw.PutString("Time for pathfind: " + strconv.Itoa(int(time.Since(startTime) / time.Millisecond)) + "ms", 0, 21)
		for x := 0; x < len(mymap); x++ {
			for y := 0; y < len(mymap[0]); y++ {
				cw.PutChar(rune(mymap[x][y]), x, y)
			}
		}
		cw.SetFgColor(cw.MAGENTA)
		c := path
		for c != nil {
			pathx, pathy := c.X, c.Y
			cw.PutChar('*', pathx, pathy)
			offx, offy := c.GetNextStepVector()
			pathx += offx
			pathy += offy
			c = c.Child
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
	cw.PutString(strconv.Itoa(width) + "x" + strconv.Itoa(height), 0, 22)
	costmap := make([][]int, width)
	for j := range costmap {
		costmap[j] = make([]int, height)
	}
	for i:=0; i<width; i++ {
		for j:=0; j<height; j++ {
			if (*strmap)[i][j] == '#' {
				costmap[i][j] = -1
			}
		}
	}
	cw.PutString(strconv.Itoa(len(costmap)) + "x" + strconv.Itoa(len(costmap[0])), 0, 23)
	cw.Flush_console()
	return &costmap
}

