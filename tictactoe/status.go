package tictactoe

// Winner returns the number that has won if there is a winner.
// Winner returns 0 if there is no winner.
func Winner(ttt TicTacToe) int {
	combos := GenerateWinCombos(3)

	for _, combo := range combos {
		values := intMap(combo, ttt.Get)
		if all(values) {
			return values[0]
		}
	}
	return 0
}

// TODO: No idea if this works for values other than 3.
func GenerateWinCombos(size int) [][]int {
	combos := make([][]int, 0)

	// horizontals
	for i := 0; i < size; i++ {
		combo := make([]int, size)
		for j := 0; j < size; j++ {
			combo[j] = j + size*i
		}
		combos = append(combos, combo)
	}

	// verticals
	for i := 0; i < size; i++ {
		combo := make([]int, size)
		for j := 0; j < size; j++ {
			combo[j] = j*size + i
		}
		combos = append(combos, combo)
	}

	// diagonals
	combo := make([]int, size)
	for j := 0; j < size; j++ {
		combo[j] = j * (size + 1)
	}
	combos = append(combos, combo)

	combo = make([]int, size)
	for j := 0; j < size; j++ {
		combo[j] = j*(size-1) + (size - 1)
	}
	return append(combos, combo)
}

// intMap is a map function over an array of ints that transform from an int to an int
func intMap(items []int, mapper func(int) int) []int {
	for i, item := range items {
		items[i] = mapper(item)
	}
	return items
}

// all returns true if all items in slice are the same
func all(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	item := nums[0]
	for i := range nums {
		if nums[i] != item {
			return false
		}
	}
	return true
}
