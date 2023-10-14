package main

import "fmt"

// 找最大值
func SelectMax(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	} else {
		max := arr[0]
		for i := 1; i < len(arr); i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}

}

// 选择排序
func SelectSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		for i := 0; i < len(arr)-1; i++ {
			max := i
			for j := i + 1; j < len(arr); j++ {
				if arr[max] < arr[j] {
					max = j
				}
			}
			if max != i {
				arr[i], arr[max] = arr[max], arr[i]
			}
		}
		return arr
	}
}

// 冒泡排序,和选择排序相比就多了很多次交换过程
func BubbleSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			needExchange := false //是否需要交换
			for j := 0; j < length-i-1; j++ {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					needExchange = true
				}
			}
			if !needExchange { //只是控制打印次数，没有性能优化
				break
			}
			fmt.Println(arr)
		}
		return arr
	}
}

// 插入排序
func InsertSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 1; i < length; i++ {
			backup := arr[i]
			j := i - 1
			for j >= 0 && backup < arr[j] {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = backup //多减了一次，得加回来
			fmt.Println(arr)
		}
		return arr
	}

}

func HeapSortOnce(arr []int, length int) []int {
	if length <= 1 {
		return arr
	} else {
		depth := length/2 - 1 //树的深度
		for i := depth; i >= 0; i-- {
			topMax := i //假定max在顶点
			leftChild := 2*i + 1
			rightChild := 2*i + 2
			if leftChild < length && arr[leftChild] > arr[topMax] {
				topMax = leftChild
			}
			if rightChild < length && arr[rightChild] > arr[topMax] {
				topMax = rightChild
			}
			if topMax != i {
				arr[i], arr[topMax] = arr[topMax], arr[i]
			}
		}
		return arr
	}

}

// 堆排序:二叉树原理,存储结构为数组，逻辑结构可当做二叉树,如父节点下标为n，其子节点下标为2*n+1和2*n+2

func HeapSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := length; i >= 1; i-- {
			lastIndex := i
			HeapSortOnce(arr, i)
			arr[lastIndex-1], arr[0] = arr[0], arr[lastIndex-1] //每次将极大值放到数组后面
		}
	}
	return arr
}

//快速排序

func QuickSort(arr []int) {

}

func main() {
	arr := []int{11, 10, 100}
	//max := SelectMax(arr)
	//fmt.Println(max)
	//fmt.Println(SelectSort(arr))
	//fmt.Println(InsertSort(arr))
	fmt.Println(HeapSort(arr))
}
