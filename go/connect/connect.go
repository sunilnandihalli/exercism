package connect


func ResultOf(board []string) (winner string, err error) {
	m := len(board)
	n := len(board[0])
	X := "OX"[1]
	O := "OX"[0]
	deltas := [3][2]int{{0, 1}, {-1, 1}, {-1, 0}}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	neighbours := func(i, j int, cur_player byte) (ch chan [2]int) {
		ch = make(chan [2]int)
		go func() {
			for p, row := range deltas {
				i1 := i + p - 1
				if i1 < 0 || i1 >= m {
					continue
				}
				for _, delta := range row {
					j1 := j + delta
					if j1 < 0 || j1 >= n || board[i1][j1] != cur_player {
						continue
					}
					ch <- [2]int{i1, j1}
				}
			}
			close(ch)
		}()
		return
	}

	goal_reached := func(l [2]int, cur_player byte) bool {
		return (cur_player == O && l[0] == m-1) || (cur_player == X && l[1] == n-1)
	}
	pop := func(locs [][2]int) (new_locs [][2]int, elem [2]int) {
		n := len(locs)
		elem = locs[n-1]
		new_locs = locs[:n-1]
		return
	}
	explore := func(i, j int, cur_player byte) bool {
		front := make([][2]int, 0, 10)
		front = append(front, [2]int{i, j})
		visited[i][j] = true
		var cur_loc [2]int
		for len(front) > 0 {
		    	front, cur_loc = pop(front)
			if goal_reached(cur_loc, cur_player) {
				return true
			}
			for nl := range neighbours(cur_loc[0], cur_loc[1], cur_player) {
				if !visited[nl[0]][nl[1]] {
					front = append(front, nl)
					visited[nl[0]][nl[1]] = true
				}
			}
		}
		return false
	}
	winner = ""
	err = nil
	for _, player := range [2]byte{O, X} {
		if player == O {
			for j := 0; j < n; j++ {
				if board[0][j] == player && !visited[0][j] && explore(0, j, player) {
					winner = string([]byte{player})
					return
				}
			}
		} else if player == X {
			for i := 0; i < m; i++ {
				if board[i][0] == player && !visited[i][0] && explore(i, 0, player) {
					winner = string([]byte{player})
					return
				}
			}
		}
	}
	return
}
