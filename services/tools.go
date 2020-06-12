package services

func isItEnough(pos *int, out *[]int) bool {
	if *pos == len(*out) {
		return true
	}
	return false
}

func sliceCheck(in *[][]int) bool {
	result := true
	for _, sl := range *in {
		if len(sl) != len((*in)[0]) {
			result = false
		}
	}
	return result
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
