package main

import (
	"fmt"
	"os"
	)

func main() {
	maze := readMaze("src/algorithm/maze/maze.in")
	start:=Point{0, 0}
	end:=Point{len(maze) - 1, len(maze[0]) - 1}
	steps,endNode := walk(maze, start, end)
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}

	for node:=endNode;node!=nil;node=node.preNode{
		fmt.Println(node.row,node.colum,node.val)
	}



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
}

func (p Point) add(dir Point) Point {
	return Point{p.row + dir.row, p.colum + dir.colum}
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
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func walk(maze [][]int, start, end Point) (steps [][]int,endNode *Node) {
	steps = make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	//points := []Point{start}
	points := []*Node{creatNode(start,0)}
	for len(points) > 0 {
		cur := points[0]
		if cur.Point == end {
			endNode=cur
			break
		}
		curSteps, _ := cur.getValue(steps)
		points = points[1:]
		for _, dir := range dirs {
			next := cur.add(dir)
			if val, ok := next.getValue(maze); !ok || val == 1 {
				continue
			}
			if nextStep, ok := next.getValue(steps); !ok || nextStep != 0{
				continue
			}
			if next == start {
				continue
			}
			val :=curSteps + 1
			steps[next.row][next.colum]  =val
			nextNode:=creatNode(next,val)
			nextNode.preNode=cur
			points = append(points, nextNode)
		}
	}
	return
}

type Node struct {
	Point
	val int
	preNode *Node
}

func creatNode(p Point,val int) *Node  {
	return &Node{Point:Point{p.row,p.colum},val:val,preNode: nil}
}






