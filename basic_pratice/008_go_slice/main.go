package main

import "fmt"

func main(){

	//slice use for hash or dynamic array

	var product = []string{"Mobile", "Laptop", "Tablet"}
	fmt.Println(product)

	marks := []int{90,20,70,80}
	marks =	append(marks, 100)

	fmt.Println(marks)
	fmt.Println("first element: ", marks[0])
	fmt.Println("last element:", marks[len(marks)-1])

	// length and capicity of slice
	//make([], len, cap)

	subjects := make([]string, 0, 3)
	fmt.Println(subjects)
	fmt.Println("length:", len(subjects))
	fmt.Println("capacity:", cap(subjects))
	subjects = append(subjects, "Math", "Science", "English")
	fmt.Println(subjects)
	fmt.Println("length:", len(subjects))
	fmt.Println("capacity:", cap(subjects))
	subjects = append(subjects, "Hindi")
	fmt.Println(subjects)
	fmt.Println("length:", len(subjects))
	fmt.Println("capacity:", cap(subjects))

	arr := [5]int{1,2,3,4,5}

	//s := arr[1:4] // [2 3 4]
	//s2 := arr[1:2] // [2]
	//s[1] = 20 // it is store refrence if you change in slice it will change in array also

	array2 := arr //[10 2 3 4 5] in case array literal it will copy value not copy reference
	array2[0] = 10 // [10 2 3 4 5] in case array literal it will copy value not copy reference


	arr3 := []int{1,2,3,4,5} // this is slice not array
	arr4 := arr3 // this is slice not array
	arr4[0] = 10 // this will change in arr3 also because it is slice and it store reference
	fmt.Println("arr3 : ", arr3) //[10 2 3 4 5] in the case of slice it will copy reference not value
	fmt.Println("arr4 : ", arr4) //[10 2 3 4 5] in the case of slice it will copy reference not value

	//fmt.Println(s)
	//fmt.Println(s2)
	fmt.Println("arr : ", arr)
	fmt.Println("array2 : ", array2)

	// coppy array
	array_a := []int {1,2,3,4,5}
	fmt.Println("array_a : ", array_a)
	array_b := array_a

	array_b[0] = 10
	fmt.Println("array_a : ", array_a)
	fmt.Println("array_b : ", array_b)

	// merge slice
	slice1 := []int{1,2,3}
	slice2 := []int{4,5,6}
	mergedSlice := append(slice1, slice2...)
	fmt.Println("mergedSlice : ", mergedSlice)	
}