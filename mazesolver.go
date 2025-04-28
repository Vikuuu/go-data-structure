package datstr

type Point struct {
	X int
	Y int
}

var dir = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func walk(maze []string, wall string, curr Point, end Point, seen [][]bool, path *[]Point) bool {
	// 1. Base case
	// off the map
	if curr.X < 0 || curr.X >= len(maze[0]) ||
		curr.Y < 0 || curr.Y >= len(maze) {
		return false
	}
	// on a wall
	if string(maze[curr.Y][curr.X]) == wall {
		return false
	}
	// at the end
	if curr.X == end.X && curr.Y == end.Y {
		*path = append(*path, curr)
		return true
	}
	// if we have seen it
	if seen[curr.Y][curr.X] {
		return false
	}

	// 3 recursion
	// pre
	seen[curr.Y][curr.X] = true
	*path = append(*path, curr)

	// recursion
	for i := 0; i < len(dir); i++ {
		x, y := dir[i][0], dir[i][1]
		if walk(maze, wall, Point{
			X: curr.X + x,
			Y: curr.Y + y,
		}, end, seen, path) {
			return true
		}
	}

	// post
	*path = (*path)[:len(*path)-1]

	return false
}

func MazeSolver(maze []string, wall string, start, end Point) []Point {
	seen := make([][]bool, len(maze))
	path := []Point{}

	for i := 0; i < len(maze); i++ {
		seen[i] = make([]bool, len(maze[i]))
	}

	walk(maze, wall, start, end, seen, &path)
	return path
}
