package main

import "fmt"

func main() {
	array1 := []int{10, 15, 2, 3}
	array2 := []int{14, -5, 21, 0}
	fmt.Println("sorted array1")
	sArray1 := sort(array1)
	fmt.Println(sArray1)
	fmt.Println("sorted array2")
	sArray2 := sort(array2)
	fmt.Println(sArray2)
	Array3 := merge(sArray1, sArray2)
	fmt.Println("Merged array is ")
	fmt.Println(Array3)
	return

}

func sort(array1 []int) (array3 []int) {
	for j := 0; j < len(array1); j++ {
		for i := 0; i < len(array1); i++ {
			if array1[i] > array1[j] {
				x := array1[j]
				array1[j] = array1[i]
				array1[i] = x
			}
		}
	}
	return array1

}

func merge(sArray1 []int, sArray2 []int) (array [8]int) {
	var mArray [8]int
	i := 0
	j := 0
	k := 0
	for i < len(sArray1) && j < len(sArray2) {
		if sArray1[i] < sArray2[j] {
			mArray[k] = sArray1[i]
			i++
		} else {
			mArray[k] = sArray2[j]
			j++
		}
		k++
	}
	if i >= len(sArray1) {
		for j < len(sArray2) {
			mArray[k] = sArray2[j]
			j++
			k++
		}
	}
	if j >= len(sArray2) {
		for i < len(sArray1) {
			mArray[k] = sArray1[i]
			i++
			k++
		}
	}

	return mArray

}
