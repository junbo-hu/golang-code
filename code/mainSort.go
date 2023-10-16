package main

// 各种排序算法
import (
	"fmt"
	"math/rand"
)

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

//快速排序，分治算法，在所有排序中速度最快

func QuickSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		n := rand.Int() % length
		splitData := arr[n]       //随机一个为基准
		low := make([]int, 0, 0)  //存储比splitData小的
		high := make([]int, 0, 0) //存储比splitData大的
		mid := make([]int, 0, 0)  //存储与splitData相等的
		mid = append(mid, splitData)
		for i := 0; i < length; i++ {
			if i == n {
				continue
			}
			if arr[i] < splitData {
				low = append(low, arr[i])
			} else if arr[i] > splitData {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high) //切割递归处理
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}

// 奇偶排序,先奇数位冒泡，再偶数位冒泡，如此循环直到数组有序

func OddEvenSort(arr []int) []int {
	length := len(arr)
	isSorted := false
	for !isSorted {
		isSorted = true
		for i := 1; i < length-1; i += 2 { //奇数位
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false //需要发生交换即还未完成
			}
		}
		fmt.Println("1:", arr)
		for i := 0; i < length-1; i += 2 { //偶数位
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false //需要发生交换即还未完成
			}

		}
		fmt.Println("0:", arr)
	}
	return arr
}

func merge(leftarr []int, rightarr []int) []int {
	leftindex := 0
	rightindex := 0
	lastarr := []int{}
	for leftindex < len(leftarr) && rightindex < len(rightarr) {
		if leftarr[leftindex] < rightarr[rightindex] {
			lastarr = append(lastarr, leftarr[leftindex])
			leftindex++
		} else if leftarr[leftindex] > rightarr[rightindex] {
			lastarr = append(lastarr, rightarr[rightindex])
			rightindex++
		} else { //相等，两边都加
			lastarr = append(lastarr, leftarr[leftindex])
			lastarr = append(lastarr, rightarr[rightindex])
			leftindex++
			rightindex++
		}
	}

	//把剩余的元素加进来
	for leftindex < len(leftarr) {
		lastarr = append(lastarr, leftarr[leftindex])
		leftindex++
	}

	for rightindex < len(rightarr) {
		lastarr = append(lastarr, rightarr[rightindex])
		rightindex++
	}
	return lastarr
}

// 归并排序

func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 { // 优化，若length<10可选用我插入排序
		return arr
	} else {
		mid := length / 2
		leftarr := MergeSort(arr[:mid])
		rightarr := MergeSort(arr[mid:])

		return merge(leftarr, rightarr) //合并
	}
}

func ShellSortStep(arr []int, start int, gap int) {
	length := len(arr)
	for i := start + gap; i < length; i += gap { // 进行插入排序
		backup := arr[i]
		j := i - gap
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = backup //多减了一次，得加回来
		fmt.Println(arr)
	}
}

// 希尔排序,步长收缩排序，可用多线程

func ShellSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		gap := length / 2
		for gap > 0 {
			for i := 0; i < gap; i++ {
				ShellSortStep(arr, i, gap)
			}
			gap /= 2
		}
		return arr
	}

}

// 桶排序或基数排序

func RadixSort(arr []int) []int {

}

func main() {
	arr := []int{11, 10, 100, 9, 34}
	//max := SelectMax(arr)
	//fmt.Println(max)
	//fmt.Println(SelectSort(arr))
	//fmt.Println(InsertSort(arr))
	//fmt.Println(HeapSort(arr))
	fmt.Println(ShellSort(arr))
}
