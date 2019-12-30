package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	maze := readMaze("src/algorithm/maze/maze.in")
	start:=Point{0, 0,0}
	end:=Point{len(maze) - 1, len(maze[0]) - 1,0}
	steps := walk(maze, start, end)
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}


	walkLine(steps,start,end)



}

func readMaze(filename string) (maze [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze = make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return
}

type Point struct {
	row, colum int
	val int
}

func (p Point) add(dir Point) Point {
	return Point{p.row + dir.row, p.colum + dir.colum,p.val+dir.val}
}

func (p Point) getValue(grid [][]int) (int, bool) {
	if p.row < 0 || p.row >= len(grid) {
		return 0, false
	}

	if p.colum < 0 || p.colum >= len(grid[p.row]) {
		return 0, false
	}

	return grid[p.row][p.colum], true
}



var dirs = [4]Point{
	{-1, 0,-1}, {0, -1,-1}, {1, 0,-1}, {0, 1,-1},
}



func walk(maze [][]int, start, end Point) (steps [][]int) {
	steps = make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}


	points := []Point{start}

	for len(points) > 0 {
		cur := points[0]
		if cur == end {
			break
		}
		curSteps, _ := cur.getValue(steps)

		points = points[1:]

		for _, dir := range dirs {
			next := cur.add(dir)
			if val, ok := next.getValue(maze); !ok || val == 1 {
				continue
			}

			nextStep, ok := next.getValue(steps)
			if !ok || nextStep != 0 {
				continue
			}

			if next == start {
				continue
			}

			steps[next.row][next.colum]  =curSteps + 1

			points = append(points, next)
		}



	}
	return
}

type pointSlice  []Point

func (p pointSlice) Len() int           { return len(p) }
func (p pointSlice) Less(i, j int) bool {
	if p[i].val == p[j].val{
		if p[i].row==p[j].row{
			return p[i].colum<p[j].colum
		}
		return p[i].row<p[j].row
	}
	return p[i].val < p[j].val


}
func (p pointSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }



func walkLine(steps [][]int,start, end Point){
	points :=make(pointSlice,0)

	for i,row :=range steps{
		for j,val:=range row{
			if steps[i][j]>0{
				points=append(points,Point{i,j,val})
			}
		}
	}
	end =points[len(points)-1]
	sort.Sort(sort.Reverse(points))

	for len(points)>0{
		cur := points[0]
		points=points[1:]
		if cur ==end{
			fmt.Println(cur)
			for _, dir := range dirs {
				//next := cur.add(dir)
				//points.


			}
		}


		//for _,v :=range points{
		//
		//}
	}



	fmt.Println(points)
}
