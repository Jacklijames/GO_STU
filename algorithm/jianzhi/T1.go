package jianzhi

func Find(target int, array [][]int ) bool{

	y := len(array) - 1
	x := 0
	for  {
		if  x >= len(array) || y < 0 {
			break
		}
		if array[x][y] < target {
			x++
		}else if array[x][y] > target {
			y--
		}else {
			return true
		}
	}

	return false
}
