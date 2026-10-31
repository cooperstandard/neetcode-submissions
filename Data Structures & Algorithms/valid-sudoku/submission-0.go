func isValidSudoku(board [][]byte) bool {
	// stored as index to val and number of times seen
	rows := make([]map[byte]int, len(board))
	cols := make([]map[byte]int, len(board))
	houses := make([]map[byte]int, len(board))

	for i, _ := range rows {
		rows[i] = make(map[byte]int)
		cols[i] = make(map[byte]int)
		houses[i] = make(map[byte]int)
	}

	for i, v := range board {
		for _, elem := range v {
			rows[i][elem]++
		}
	}

	for _, v := range board {
		for j, elem := range v {
			cols[j][elem]++
		}
	}

	// traverse each house
	// 0, 1, 2
	// 3, 4, 5
	// 6, 7, 8
	// (i / 3) + ((j/3) * 3) 
	for i, v := range board {
		for j, elem := range v {
			houses[(i/3)+((j/3)*3)][elem]++
		}
	}

	for i, _ := range rows {
		for k, v := range rows[i] {
			if k == '.' {
				continue
			}
			if v > 1 {
				return false
			}
		}
		for k, v := range cols[i] {
			if k == '.' {
				continue
			}
			if v > 1 {
				return false
			}
		}
		for k, v := range houses[i] {
			if k == '.' {
				continue
			}
			if v > 1 {
				return false
			}
		}
	}
	return true



}
