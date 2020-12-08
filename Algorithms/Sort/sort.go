package sort

import(
	"fmt"
)

func BubbleSort(input []int) (res []int, err error){
	// fmt.Println("BubbleSort input:", input)
	count := 0
	for i := 0; i < len(input); i++{
		for j := len(input) - 1; j -i > 0; j--{
			if input[j] < input[j-1]{
				tmp := input[j]
				input[j] = input[j-1]
				input[j-1] = tmp				
			}
			count++
		}
	}
	fmt.Println("calculate counter:", count, "input length:", len(input))
	res = input
	// fmt.Println("BubbleSort output:", res)
	err = nil
	return
}

func InsertionSort(input []int) (res []int, err error){
	count := 0

	fmt.Println("calculate counter:", count, "input length:", len(input))
	res = input
	err = nil
	return
}

func SelectionSort(input []int) (res []int, err error){
	count := 0

	fmt.Println("calculate counter:", count, "input length:", len(input))
	res = input
	err = nil
	return
}

func HeapSort(input []int) (res []int, err error){
	count := 0

	fmt.Println("calculate counter:", count, "input length:", len(input))
	res = input
	err = nil
	return
}

func MergeSort(input []int) (res []int, err error){
	count := 0

	fmt.Println("calculate counter:", count, "input length:", len(input))
	res = input
	err = nil
	return
}

func QuickSort(input []int) (res []int, err error){
	fmt.Println("QuickSort")

	res = input
	err = nil
	return
}