func spiralOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	res := []int{}

	srow, scol, erow, ecol := 0, 0, m-1, n-1

	for srow <= erow && scol <= ecol {
		for i := scol; i <= ecol; i++ {
			res = append(res, mat[srow][i])
		}

		for i := srow + 1; i <= erow; i++ {
			res = append(res, mat[i][ecol])
		}

		if srow < erow {
			for i := ecol - 1; i >= scol; i-- {
				res = append(res, mat[erow][i])
			}
		}

		if scol < ecol {
			for i := erow - 1; i >= srow+1; i-- {
				res = append(res, mat[i][srow])
			}
		}
		srow++
		scol++
		erow--
		ecol--
	}
	return res
}