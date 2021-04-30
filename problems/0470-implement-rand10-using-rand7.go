package problems

func rand10() int {
	var i, j, idx int
	for {
		i = rand7()
		j = rand7()
		idx = i + (j-1)*7
		if idx <= 40 {
			return 1 + (idx-1)%10
		}
	}
}
