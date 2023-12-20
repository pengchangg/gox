package basic

import "math/rand"

func Array(n int) []int {
	return make([]int, n)
}

// ArrayRandomAccess 随机访问数组元素
func ArrayRandomAccess(array []int) int {
	return array[rand.Intn(len(array))]
}

// ArrayFind 查找数组元素
func ArrayFind(array []int, target int) int {
	for i, v := range array {
		if v == target {
			return i
		}
	}
	return -1
}

// ArrayInsert 插入元素
func ArrayInsert(array []int, index, value int) []int {
	array = append(array, 0)
	copy(array[index+1:], array[index:])
	array[index] = value
	return array
}

// ArrayDelete 删除元素
func ArrayDelete(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

// ArrayExpand 数组扩容
func ArrayExpand(array []int, n int) []int {
	return append(array, make([]int, n)...)
}
