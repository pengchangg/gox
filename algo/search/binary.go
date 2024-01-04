package search

// BinarySearch 二分查找实现 带注释
func BinarySearch(array []int, target int) int {
	// 定义左右边界
	left, right := 0, len(array)-1
	// 循环条件
	for left <= right {
		// 中间值
		mid := left + (right-left)/2
		// 目标值在左边
		if target < array[mid] {
			right = mid - 1
			// 目标值在右边
		} else if target > array[mid] {
			left = mid + 1
			// 找到目标值
		} else {
			return mid
		}
	}
	// 未找到目标值
	return -1
}
