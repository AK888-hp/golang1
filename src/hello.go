package main

import (
	"fmt"
)
func main() {
	//age:=20
	// if age>=18 {
	// 	fmt.Println("Adult")
	// } else {
	// 	fmt.Println("Minor")
	// }

	// if x:=10;x>5{
	// 	fmt.Println("x is greater than 5")
	// }
// 	for i:=0;i<5;i++{
// 		fmt.Println(i)
// 	}
// 	fmt.Println()

// 	j:=0
// 	for j<3 {
// 		fmt.Println(j)
// 		j++
// 	}
// 	numbers:=[]int{1,2,3,4,5,6,7,8,9,10}
// 	for index,value := range numbers {
// 		fmt.Printf("Index: %d, Value: %d\n", index, value)
// 	}
// 	 k:=0
// 	 for k>=0 {
// 		fmt.Println(k)
// 		k++
// 		if k==10 {
// 			break
// 		}
// 	 }

// slice:=[]int{1,2,3,4,}
// fmt.Println(slice)
// fmt.Println(slice[1:3])
// fmt.Println(len(slice))
// fmt.Println(cap(slice))
// fmt.Println(slice[1:4],len(slice[1:4]),cap(slice[1:4]))

// str:="Hello, Go!"
// fmt.Println(len(str))
// fmt.Println(str[0])
// fmt.Println(str[0:5])
// fmt.Println(strings.Contains(str,"GO"))	
// fmt.Println(strings.ToUpper(str))
// fmt.Println(strings.ToUpper(str))

	//Creating a map
	myMap := make(map[string]int)

	//Accessing values
	myMap["apple"] = 1
	myMap["banana"] = 2

	//
	fmt.Println(myMap["apple"])

	//Checking if map exists
	value, ok := myMap["grape"]

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Key not found")
	}

	//Deleting a map
	delete(myMap, "banana")
	fmt.Println(myMap)

	for key, value := range myMap {
		fmt.Printf("Key :%s, Value :%d\n", key, value)
	}
	
}