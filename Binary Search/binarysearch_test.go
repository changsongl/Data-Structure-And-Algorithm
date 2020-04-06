package Binary_Search

import "testing"

func TestSearch(t *testing.T) {
	bt := []int{1, 3, 6, 8, 11}
	testCase := map[int]int{1: 0, 6: 2, 8: 3, 9: -1, 12: -1}

	for num, index := range testCase {
		result := Search(bt, num)
		if result != index {
			t.Errorf(
				"Binary Search Failed: Finding number %d. Expect index %d, result %d",
				num, index, result)
		}
	}
}
