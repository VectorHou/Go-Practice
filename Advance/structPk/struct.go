package structPk

import (
	"fmt"
	"errors"
)

type DataType struct{
	num uint
	name string
	slices []int
	maps map[string]int
	arrays [5]int
}

func Structfunc(){// The first character of the function name must be upper-character， otherwise it can not be called by other package
	datatype := DataType{
		5,
		"Slice",
		[]int{1,2,3},
		map[string]int{
			"tom":12,
			"jack":13,
		},
		[5]int{1,2,3,4,5},
	}
	//DataType.num = 7
	datatype.GetNum()
	datatype.AlterNum(66)
	
	datatype.GetName()
	datatype.AlterName("")
	datatype.AlterName("datatype")	

	datatype.GetSlices()
	datatype.AlterSlices()

	datatype.GetMaps()
	datatype.AlterMaps()

	datatype.GetArrays()
	datatype.AlterArrays()
	fmt.Println(datatype)
}

func (dt *DataType)AlterNum(_num uint)(ret int, err error){
	dt.num = _num
	ret = 0
	err = nil
	fmt.Println("DataType::AlterNum:", _num)
	return
}

func (dt *DataType)GetNum()(_num uint){
	_num = dt.num
	fmt.Println("DataType::GetNum:", _num)
	return
}

func (dt *DataType)AlterName(_name string)(ret int, err error){
	if len(_name) > 0{
		dt.name = _name
		ret = 0
		err = nil
		fmt.Println("DataType::AlterName:", _name)
	}else{
		ret = -1
		err = errors.New("the len of name is null!\n")
		fmt.Println("DataType::AlterName:", _name, "failed!")
	}
	return
}

func (dt *DataType)GetName()(string){
	fmt.Println("DataType::GetName return", dt.name)
	return dt.name
}

func (dt *DataType)AlterSlices()(int, error){

	fmt.Println("DataType::AlterSlices:", dt.slices)
	return 0,nil
}

func (dt *DataType)GetSlices()([]int){
	fmt.Println("DataType::GetSlices return", dt.slices)
	return dt.slices
}

func (dt *DataType)AlterMaps()(int, error){

	fmt.Println("DataType::AlterMaps:", dt.maps)
	return 0,nil
}

func (dt *DataType)GetMaps()([]int){
	fmt.Println("DataType::GetMaps return", dt.maps)
	return dt.slices
}

func (dt *DataType)AlterArrays()(int, error){

	fmt.Println("DataType::AlterArrays:", dt.arrays)
	return 0,nil
}

func (dt *DataType)GetArrays()([]int){
	fmt.Println("DataType::GetArrays return", dt.arrays)
	return dt.slices
}