package main

import "fmt"

// 删除数组中指定元素并返回删除后数组长度
func removeElement(arr3 []int, val int) int {
	left := 0
	len := len(arr3)

	for right := 0; right < len; right++ {
		if arr3[right] != val {
			arr3[left] = arr3[right]
			left++
		}
	}
	return left
}

func mergeArr(arr1 []int, arr2 []int) []int {
	index1 := 0
	index2 := 0
	len1 := len(arr1)
	len2 := len(arr2)
	finalArr := []int{}
	for index1 < len1 && index2 < len2 {
		if arr1[index1] < arr2[index2] {
			finalArr = append(finalArr, arr1[index1])
			index1++
		} else if arr1[index1] > arr2[index2] {
			finalArr = append(finalArr, arr2[index2])
			index2++
		} else {
			finalArr = append(finalArr, arr1[index1])
			finalArr = append(finalArr, arr2[index2])
			index1++
			index2++
		}
	}
	for index1 < len1 {
		finalArr = append(finalArr, arr1[index1])
		index1++
	}
	for index2 < len2 {
		finalArr = append(finalArr, arr2[index2])
		index2++
	}

	return finalArr
}

func main() {
	arr1 := []int{1, 3, 4, 5, 7, 9, 23}
	arr2 := []int{2, 3, 5, 7, 8, 10, 33, 44, 55, 66}

	finalArr := mergeArr(arr1, arr2)

	fmt.Println(finalArr)

	arr3 := []int{1, 1, 1, 2, 2, 4, 6, 33, 33, 33}

	//lenArr3 := len(arr3)
	//len := removeElement(arr3, 1)
	//arr3 = arr3[:lenArr3-(lenArr3-len)]
	//
	//fmt.Println(len)
	//fmt.Println(arr3)

	lenArr3 := len(arr3)
	len := removeDuplicates2(arr3)
	arr3 = arr3[:lenArr3-(lenArr3-len)]

	fmt.Println(len)
	fmt.Println(arr3)

}

// 删除有序数组中的重复项
func removeDuplicates(arr []int) int {
	len := len(arr)
	if len == 0 || arr == nil {
		return 0
	}

	index1 := 0
	index2 := 1

	for index2 < len {
		if arr[index1] != arr[index2] {
			arr[index1+1] = arr[index2]
			index1++
		}
		index2++
	}

	return index1 + 1
}

// 删除有序数组中的重复项2
func removeDuplicates2(arr []int) int {
	len := len(arr)
	if len <= 2 {
		return len
	}

	slow := 2
	fast := 2

	for fast < len {
		if arr[slow-2] != arr[fast] {
			arr[slow] = arr[fast]
			slow++
		}
		fast++
	}

	return slow
}