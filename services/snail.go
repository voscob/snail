package services

var verticalLength, horizontalLength int

func Untwist(in *[][]int) ([]int, error) {
	if !sliceCheck(in) {
		return nil, ErrorIncorrectSlice
	}

	verticalLength = len(*in)
	horizontalLength = len((*in)[0])

	out := make([]int, horizontalLength*verticalLength)
	pos := 0

	for i := 0; ; i++ {
		leftToRight(in, i, &pos, &out)
		if isItEnough(&pos, &out) {
			break
		}

		topToBottom(in, i, &pos, &out)
		if isItEnough(&pos, &out) {
			break
		}

		rightToLeft(in, i, &pos, &out)
		if isItEnough(&pos, &out) {
			break
		}

		bottomToTop(in, i, &pos, &out)
		if isItEnough(&pos, &out) {
			break
		}
	}

	return out, nil
}

func leftToRight(in *[][]int, lap int, pos *int, out *[]int) {
	for i := lap; i < horizontalLength-lap; i++ {
		(*out)[*pos] = (*in)[lap][i]
		*pos++
	}
}

func topToBottom(in *[][]int, lap int, pos *int, out *[]int) {
	for i := lap + 1; i < verticalLength-lap; i++ {
		(*out)[*pos] = (*in)[i][horizontalLength-lap-1]
		*pos++
	}
}

func rightToLeft(in *[][]int, lap int, pos *int, out *[]int) {
	for i := horizontalLength - lap - 2; i >= lap; i-- {
		(*out)[*pos] = (*in)[verticalLength-lap-1][i]
		*pos++
	}
}

func bottomToTop(in *[][]int, lap int, pos *int, out *[]int) {
	for i := verticalLength - lap - 2; i > lap; i-- {
		(*out)[*pos] = (*in)[i][lap]
		*pos++
	}
}
