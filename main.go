package main

import (
	"fmt"
	"sdlManager-mysql/initialize"
	"sdlManager-mysql/router"
)

func init() {
	initialize.Init()
}

func main() {
	fmt.Println("gorm coding")
	engine := router.GetEngine()
	if err := engine.Run(":8060"); err != nil {
		panic(err)
	}

}

// 调整堆，使得以rootIndex为根的子树成为一个大顶堆
func heapify(arr []int, n int, rootIndex int) {
	largest := rootIndex
	left := 2*rootIndex + 1
	right := 2*rootIndex + 2

	// 找到rootIndex、left和right中的最大值索引
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值不是rootIndex，则交换并继续调整堆
	if largest != rootIndex {
		arr[rootIndex], arr[largest] = arr[largest], arr[rootIndex]
		heapify(arr, n, largest)
	}
}

// 堆排序算法
func heapSort(arr []int) {
	n := len(arr)

	// 构建初始大顶堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 从堆中取出元素并进行堆排序
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // 将最大元素放到数组末尾
		heapify(arr, i, 0)              // 调整剩余的堆
	}
}
