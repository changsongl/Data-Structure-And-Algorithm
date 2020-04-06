package Binary_Search

// Base on binary tree slice to find the number.
// return -1 not exist, otherwise return index
func Search(bt []int, num int) int{
	low, high := 0, len(bt) - 1
	for low <= high{
		mid := (low + high) / 2
		switch true{

		case bt[mid] == num:
			return mid
		case bt[mid] > num:
			high = mid - 1
		default:
			low = mid + 1
		}
	}
	return -1
}
