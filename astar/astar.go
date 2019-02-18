package astar

type direction struct {
	x, y int
}

const (
	DIAGONAL_COST = 14
	STRAIGHT_COST = 10
)

type cell struct {
	X, Y            int
	G, H            int
	costToMoveThere int
	parent          *cell
}

func (c *cell) getF() int {
	return c.G + c.H
}

func (c *cell) setG(inc int) {
	if c.parent != nil {
		c.G = c.parent.G + inc
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattansHeuristic(fromx, fromy, tox, toy int) int {
	return 10 * (abs(tox-fromx) + abs(toy-fromy))
}

func getIndexOfCellWithLowestF(openList []*cell) int {
	cheapestCellIndex := 0
	for i, c := range openList {
		if c.getF() < openList[cheapestCellIndex].getF() {
			cheapestCellIndex = i
		}
	}
	return cheapestCellIndex
}

func (c *cell) getPathToCell() *[]*cell {
	path := make([]*cell, 0)
	curcell := c
	for curcell != nil {
		path = append(path, curcell)
		curcell = curcell.parent
	}
	return &path
}

func FindPath(costMap *[][]int, fromx, fromy, tox, toy int) *[]*cell {
	openList := make([]*cell, 0)
	closedList := make([]*cell, 0)
	var currentCell *cell
	total_steps := 0
	targetReached := false
	// step 1
	origin := &cell{X: fromx, Y: fromy, costToMoveThere: 0, H: manhattansHeuristic(fromx, fromy, tox, toy)}
	openList = append(openList, origin)
	// step 2
	for !targetReached {
		// sub-step 2a:
		currentCellIndex := getIndexOfCellWithLowestF(openList)
		currentCell = openList[currentCellIndex]
		// sub-step 2b:
		closedList = append(closedList, currentCell)
		openList = append(openList[:currentCellIndex], openList[currentCellIndex+1:]...) // this friggin' magic removes currentCellIndex'th element from openList
		//sub-step 2c:
		analyzeNeighbors(currentCell, &openList, &closedList, costMap, tox, toy)
		//sub-step 2d:
		total_steps += 1
		if getCellWithCoordsFromList(&openList, tox, toy) != nil {
			return currentCell.getPathToCell()
		}
		if len(openList) == 0 {
			return &[]*cell {}
		}
	}
	return &[]*cell {}
}

func analyzeNeighbors(curCell *cell, openlist *[]*cell, closedlist *[]*cell, costMap *[][]int, targetX, targetY int) {
	cost := 0
	cx, cy := curCell.X, curCell.Y
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			x, y := cx+i, cy+j
			if areCoordsValidForCostMap(x, y, costMap) {
				// TODO: maybe include target cell even if it is not passable?
				if (*costMap)[x][y] == -1 || getCellWithCoordsFromList(closedlist, x, y) != nil { // cell is impassable or is in closed list
					continue // ignore it
				}
				// TODO: add a flag for skipping diagonally lying cells
				if (i * j) != 0 { // the cell under consideration is lying diagonally
					cost = DIAGONAL_COST
				} else {
					cost = STRAIGHT_COST
				}
				curNeighbor := getCellWithCoordsFromList(openlist, x, y)
				if curNeighbor != nil {
					if curNeighbor.G > curCell.G + cost {
						curNeighbor.parent = curCell
						curNeighbor.setG(cost)
					}
				} else {
					curNeighbor = &cell{X: x, Y: y, parent:curCell, H:manhattansHeuristic(x, y, targetX, targetY)}
					curNeighbor.setG(cost)
					*openlist = append(*openlist, curNeighbor)
				}
			}
		}
	}
}

func getCellWithCoordsFromList(list *[]*cell, x, y int) *cell {
	for _, c := range *list {
		if c.X == x && c.Y == y {
			return c
		}
	}
	return nil
}

func areCoordsValidForCostMap(x, y int, costMap *[][]int) bool {
	return x >= 0 && y >= 0 && (x < len(*costMap)) && (y < len((*costMap)[0]))
}

// АЛГОРИТМ А*:
// Замечание: F = G+H, где G - цена пути ИЗ стартовой точки, H - эвристическая оценка пути ДО цели.
// По "методу Манхэттена": H = 10 * (_abs(targetX - startX) + _abs(targetY-startY))
//  1) Добавляем стартовую клетку в открытый список.
//  2) Повторяем следующее:
//  a) Ищем в открытом списке клетку с наименьшей стоимостью F. Делаем ее текущей клеткой.
//  b) Помещаем ее в закрытый список. (И удаляем с открытого)
//  c) Для каждой из соседних 8-ми клеток ...
//  	Если клетка непроходимая или она находится в закрытом списке, игнорируем ее. В противном случае делаем следующее.
//  	Если клетка еще не в открытом списке, то добавляем ее туда. Делаем текущую клетку родительской для это клетки. Расчитываем стоимости F, G и H клетки.
//  	Если клетка уже в открытом списке, то проверяем, не дешевле ли будет путь через эту клетку. Для сравнения используем стоимость G.
//  	Более низкая стоимость G указывает на то, что путь будет дешевле. Эсли это так, то меняем родителя клетки на текущую клетку и пересчитываем для нее стоимости G и F. Если вы сортируете открытый список по стоимости F, то вам надо отсортировать свесь список в соответствии с изменениями.
//  d) Останавливаемся если:
//  	Добавили целевую клетку в открытый список, в этом случае путь найден.
// 		Или открытый список пуст и мы не дошли до целевой клетки. В этом случае путь отсутствует.
//  3) Сохраняем путь. Двигаясь назад от целевой точки, проходя от каждой точки к ее родителю до тех пор, пока не дойдем до стартовой точки. Это и будет наш путь.
