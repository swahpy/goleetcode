package numberofsteps

func numberOfSteps(num int) int {
	steps := 0
	for num != 0 {
		if num%2 == 0 {
			num = num >> 1
		} else {
			num--
		}
		steps++
	}
	return steps
}
