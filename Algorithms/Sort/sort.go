package sort

import(
	"fmt"
	"errors"
)

func BubbleSort(input []int) (res []int, err error){
	// fmt.Println("BubbleSort input:", input)
	fmt.Println("input length:", len(input))
	if len(input) <= 0{
		err = errors.New("the size of input can not be sort!")
		return
	}
	count, execount := 0,0
	for i := 0; i < len(input); i++{
		for j := len(input) - 1; j -i > 0; j--{
			count++
			if input[j] < input[j-1]{
				execount++
				tmp := input[j]
				input[j] = input[j-1]
				input[j-1] = tmp				
			}			
		}
	}
	fmt.Println("loop counter:", count, "swap counter:", execount)
	res = input
	// fmt.Println("BubbleSort output:", res)
	err = nil
	return
}

func InsertionSort(input []int) (res []int, err error){
	fmt.Println("input length:", len(input))
	if len(input) <= 0{
		err = errors.New("the size of input can not be sort!")
		return
	}
	count, execount := 0,0
	for i := 0; i < len(input); i++{
		index := i
		for j := i; j >= 0; j--{
			count++	
			// fmt.Println(index, j)		
			if input[index] < input[j]{		
				execount++		
				tmp := input[j]
				input[j] = input[index]
				input[index] = tmp
				index = j
				// fmt.Println(input)
			}			
		}
	}
	fmt.Println("loop counter:", count, "swap counter:", execount)
	res = input
	err = nil
	return
}

func SelectionSort(input []int) (res []int, err error){
	fmt.Println("input length:", len(input))
	if len(input) <= 0{
		err = errors.New("the size of input can not be sort!")
		return
	}
	count, execount := 0,0
	for i := 0; i < len(input); i++{
		minindex := i
		for j := i; j < len(input); j++{
			count++
			if input[minindex] > input[j]{
				execount++
				minindex = j
			}
		}
		tmp := input[minindex]	
		input[minindex] = input[i]
		input[i] = tmp
	}
	res = input
	fmt.Println("loop counter:", count, "swap counter:", execount)  
	err = nil
	return
}

func HeapSort(input []int) (res []int, err error){
	fmt.Println("input length:", len(input))
	if len(input) <= 0{
		err = errors.New("the size of input can not be sort!")
		return
	}
	count, execount := 0,0

	fmt.Println("loop counter:", count, "swap counter:", execount)
	res = input
	err = nil
	return
}

func MergeSort(input []int) (res []int, err error){
	fmt.Println("input length:", len(input))
	if len(input) <= 0{
		err = errors.New("the size of input can not be sort!")
		return
	}
	count, execount := 0,0

	fmt.Println("loop counter:", count, "swap counter:", execount)
	res = input
	err = nil
	return
}

func QuickSort(input []int) (res []int, err error){
	fmt.Println("input length:", len(input))
	if len(input) <= 0{
		err = errors.New("the size of input can not be sort!")
		return
	}
	count, execount := 0,0

	fmt.Println("loop counter:", count, "swap counter:", execount)
	res = input
	err = nil
	return
}