package problem0073

func SetZeroes(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	marks := make([]int, m+n)

    for i := range marks {
        marks[i]=-1
    }

    for i, row := range matrix {
        for j, num := range row {
            if num == 0 {
                marks[i]=1
                marks[m+j]=1
            }
        }
    }

    for i, mark := range marks[:m] {
        if mark != -1 {
            matrix[i] = make([]int,n)
        }
    }

    for i, mark := range marks[m:] {
        if mark != -1 {
            for _, row := range matrix {
                row[i]=0
            }
        }
    }

	return matrix
}