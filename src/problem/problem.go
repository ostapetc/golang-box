package fairrations

func FairRations(subjects []int32) int32 {
	var result int32

	posible := changePosible(subjects)

	for posible == 1{
		changeSubjects(subjects)
		result += 2
		posible = changePosible(subjects)

		if posible == -1 {
			return 0
		}
	}

	return result
}

func changePosible(subjects []int32) int {
	for i, value := range subjects {
		if value % 2 != 0 {
			if i == len(subjects) - 2 && subjects[i + 1] % 2 == 0 {
				return -1
			} else if i == len(subjects) - 1 && subjects[i - 1] % 2 == 0 {
				return -1
			}

			return 1
		}
	}

	return 0
}

func changeSubjects(subjects []int32) {
	for i, value := range subjects {
		if value % 2 != 0 {
			subjects[i]++

			if (i + 1 == len(subjects)) { //last {
				subjects[i - 1]++
			} else {
				subjects[i + 1]++
			}

			break
		}
	}
}