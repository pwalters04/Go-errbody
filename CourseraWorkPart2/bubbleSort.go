package main

import (
	"fmt"
)

func main (){
  num := []int{ 4,2,6}
  sorted := BubbleSort(num)
  for index,_ := range sorted{
   fmt.Println(sorted[index])
  }

}

func BubbleSort(unSortedSlice []int)[]int {
	swap(unSortedSlice)
	return unSortedSlice
}
func swap( numSlice []int ){
	 sliceLength := len(numSlice)
	 var temp int
	for position := 0; position < sliceLength - 1; position ++ {
		if numSlice[position] > numSlice[position +1]{
         temp =numSlice[position]
		 numSlice[position] = numSlice[position +1]
		 numSlice[position+1] = temp
		}
	}
}