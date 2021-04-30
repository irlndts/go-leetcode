package problems

// https://leetcode.com/problems/keys-and-rooms/

func canVisitAllRooms(rooms [][]int) bool {
	visited = make([]bool, len(rooms))

	path(0, rooms)

	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}

var visited []bool

func path(room int, rooms [][]int) {
	visited[room] = true
	for _, r := range rooms[room] {
		if !visited[r] {
			path(r, rooms)
		}
	}
}
