//广度优先走迷宫算法

package main

import (
	"fmt"
	"os"
)

//读取迷宫文件的数据，返回二维切片
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	//读取行数和列数
	fmt.Fscanf(file, "%d %d", &row, &col)
	//slice的元素是其他slice，元素类型是[]int
	maze := make([][]int, row) //定义row个slice组成的slice
	//遍历每一行
	for i := range maze {
		//每一行又是一个slice，slice的元素个数还是col个
		maze[i] = make([]int, col)
		//遍历每一列
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze

}

//定义点的结构
type point struct {
	i, j int
}

//定义探索的4个方向：上，左、下、右
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}
func (p point) at(grid [][]int) (int, bool) {

	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

//走迷宫，返回
//start：起始点；end：目标点
func walk(maze [][]int, start, end point) [][]int {
	//定一个二维切片，其中的每个元素表示从start走了多少步才走到这一格
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	//将起点加入到队列中
	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		//如果发现到了终点就退出
		if cur == end {
			break
		}

		//探测四个方向
		for _, dir := range dirs {
			//下一个结点的坐标=当前结点+方向
			next := cur.add(dir)
			//下一个结点是空白地方，可以走过去才可以探索
			//下一个结点的值有，说明曾经到达过，就不能走
			//如果回到起点，也不可以走

			//尝试探索下一个结点
			val, ok := next.at(maze)
			//如果撞墙，剔除不能走的
			if !ok || val == 1 {
				continue
			}
			val, ok = next.at(steps)
			//如果已经走过了
			if !ok || val != 0 {
				continue
			}
			//如果回到了原点
			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			//填入steps
			steps[next.i][next.j] = curSteps + 1
			//将点填入到队列中
			Q = append(Q, next)
		}
	}

	return steps
}
func main() {
	//读取文件，注意在windows下，maze.in的换行符要换成LF，否则会出现末尾读入数字0
	maze := readMaze("maze/maze.in")
	//输出读取的文件值，验证是否准确
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
	//走迷宫，传入起始点和目标点
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
