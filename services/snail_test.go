package services

import (
	"testing"
)

func TestUntwist(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Name     string
		Mock     [][]int
		Expected []int
		Err      error
	}{
		{
			Name: "Wrong slice",
			Mock: [][]int{
				{1, 2, 3, 4, 5, 1},
				{9, 1, 1, 1, 8},
				{9, 1, 1, 1, 1, 7},
			},
			Expected: []int{},
			Err:      ErrorIncorrectSlice,
		},
		{
			Name: "Empty slice",
			Mock: [][]int{
				{},
			},
			Expected: []int{},
			Err:      nil,
		},
		{
			Name: "Super simple slice",
			Mock: [][]int{
				{1},
			},
			Expected: []int{1},
			Err:      nil,
		},
		{
			Name: "Just a simple slice",
			Mock: [][]int{
				{1, 2},
				{4, 3},
			},
			Expected: []int{1, 2, 3, 4},
			Err:      nil,
		},
		{
			Name: "Just a slice",
			Mock: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
			Expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			Err:      nil,
		},
		{
			Name: "Medium slice",
			Mock: [][]int{
				{1, 2, 3, 4, 5},
				{16, 17, 18, 19, 6},
				{15, 24, 25, 20, 7},
				{14, 23, 22, 21, 8},
				{13, 12, 11, 10, 9},
			},
			Expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
			Err:      nil,
		},
		{
			Name: "Vertical slice",
			Mock: [][]int{
				{1, 2, 3, 4, 5},
				{20, 21, 22, 23, 6},
				{19, 32, 33, 24, 7},
				{18, 31, 34, 25, 8},
				{17, 30, 35, 26, 9},
				{16, 29, 28, 27, 10},
				{15, 14, 13, 12, 11},
			},
			Expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35},
			Err:      nil,
		},
		{
			Name: "Horizontal slice",
			Mock: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8},
				{22, 23, 24, 25, 26, 27, 28, 9},
				{21, 36, 37, 38, 39, 40, 29, 10},
				{20, 35, 34, 33, 32, 31, 30, 11},
				{19, 18, 17, 16, 15, 14, 13, 12},
			},
			Expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40},
			Err:      nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			got, err := Untwist(&tt.Mock)

			if err != tt.Err {
				t.Error("Err:", err, "expected:", tt.Err)
			}

			if !sliceEqual(got, tt.Expected) {
				t.Error("Got:", got, "expected:", tt.Expected)
			}
		})
	}
}
