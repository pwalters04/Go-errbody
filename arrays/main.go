package main

import(
	"fmt"
	"math"
	"sort"
)
type DVD struct {
	Title string
	Year int
	Actor string
}
func main(){
	// 0 - (N-1)
 	var dvdCollection [15] *DVD
	avengersDVD := &DVD{
		Title: "The Avengers",
		Year:  2012,
		Actor: "Joss Whedon",
	}



	incrediblesDVD := &DVD{
		Title: "The Incredibles",
		Year: 2004,
		Actor: "Brad Bird",
	}

	findingDoryDVD := &DVD{
		Title: "Finding Dory",
		Year: 2016,
		Actor: "Andrew Stanton",
	}

	//Insert An Element
	dvdCollection[14] = avengersDVD
	dvdCollection[3] = incrediblesDVD
	dvdCollection[2] = findingDoryDVD

	//var intCollection [10]int
	//var intSq int
	//
	//for i:=0 ; i < 10 ; i++ {
	//	intSq = (i+1)*(i+1)
	//	 intCollection[i] = intSq
	//	 fmt.Println(intCollection[i])
	//}

	//x := []int{1}
	//
	//fmt.Println(findMaxConsecutiveOnes(x))


	//y := []int{555,901,482,1771}
	//fmt.Println(findNumbers(y))
	z := []int {-7,-3,2,3,11}
	//newZ := SortedSquares(z)
	sorted := BubbleSort(z)
	for i:=range sorted{
		fmt.Println("array: ",sorted[i])
	}

}


func findMaxConsecutiveOnes(nums []int) int {
	var countOnes int
	var maxCount int

	for i:=0 ; i <= len(nums)-1 ; i++{
		if nums[i] == 1{
			countOnes++
		}else if nums[i]<1{
			countOnes = 0
		}

		if maxCount < countOnes{
			maxCount = countOnes
		}

	}
	return maxCount
}
/*
divide by 10 until 0
count how many times we have divided by 10
  if its count % 2 = 0
      its an even number of digits

*/
//Find Numbers with Even Number of Digits
func findNumbers(nums []int) int {
	var evenCount int
	var digitCount int

	for i:=0 ; i <=len(nums)-1 ;i++{
		   digitCount =DivdeByTen(nums[i])
		if digitCount % 2 ==0{
			evenCount ++
		}
	}

return evenCount
}

func DivdeByTen( nums int) int {
	 slicing  := 0
	for nums !=0{
		nums =nums/10
		slicing ++
	}
 return slicing
}

func SortedSquares(A []int) []int {
		var sqint float64
		var val float64
	for i:= 0 ; i<= len(A)-1;i++{
		val = float64(A[i])
		sqint= math.Pow(val,2)
		A[i] =int(sqint)
	}
	sort.Ints(A)
	return A
}

func BubbleSort (arr []int) []int{
	 i:=0
	for i<=len(arr)-1{
		if arr[i] > arr[i+1]{
			arr[i],arr[i+1]=Swap(arr[i],arr[i+1])
		}
	}
	return arr
}

func Swap (item1 int, item2 int) (int,int ){
	var temp int

	temp = item1
	item1 = item2
	item2 = temp

	return item1,item2
}

func MilitaryTime (arr []int) string{
 return ""
}

func duplicateZeros(arr []int) {
	var numberZeros int
	var arrLen int = len(arr)
	for _, val := range arr{
		if val == 0{
			numberZeros ++
		}
	}

	i,j := arrLen-1 ,arrLen+numberZeros-1
	for ; i>=0 && j>=0 ; i,j = i-1 ,j-1{
		if arr[i] != 0{
			if j< arrLen{
				arr[j]=arr[i]
			}
		}else{
			if j<arrLen{
				arr[j]=0
			}
			j--
			if j<arrLen{
				arr[j]=0
			}
		}
	}

}