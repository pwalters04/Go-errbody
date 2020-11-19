package main

/*
Given an HTML string, parse the string for HTML tags and return which tags you find,
along with their count.
input: "<html><body>hello world<button></button><p>hi</p></body></html>"
output: { completed tag (open/ closed)
          uncompleted tag (open or closed) }
*/

/*
 array: string
 Search String: "<" char/word > return char or word" , -> array  or map
Char / Words:  count ++ each key {map} or  array { count value }
return: map
print map keys / values

*/

//  myHTMLString := "<html><body>hello world<button></button><p>hi</p></body></html>"
// s := strings.Split(myHTMLString, "<")
// var  wordsArray[] string
// for i,_ := range s {
// 	if s[i] == ">"{
// 		wordsArray = append(wordsArray, s[i-1])
// 	}
// }
// fmt.Println(wordsArray)
func main() {



	
}
// NumInList will return true if num is in the
// list slice and false otherwise
func NumInList( list []int,num int)bool{

	if list == nil{
		return false
	}
	for _,value := range list{
		if value == num{
			return true
		}
	}
	return false
}

// sum all nums

func SumNums( numbers []int)int{
	sum := 0

	if numbers == nil{
		return 0
	}
	for _,value := range numbers{
		sum += value
	}
	return sum
}
/*
	 letter := "AAAAAAdfgsrgg"
	 count := 0

	 for _,k := range letter{
	 	if unicode.IsUpper(k){
	 		count++
		}
	 }
 fmt.Println("Cap Letters: ", count)
 */

//x := []int{3, 5, -4, 8, 11, 1, -1, 6}
//y:= []int{5,8,1,6}
//fmt.Println(IsValidSubsequence(x,y))


//func IsValidSubsequence(array []int, sequence []int) bool {
//	var checkArrayLength int
//	var sequenceLength= len(sequence)
//	var arrayLength = len(array)
//	var checkSequenceLength int
//
//	matchCount:=0
//	j:=0
//	i:=0
//
//	for {
//		if sequence[i] == array[j]{
//			checkSequenceLength++
//			i++
//			matchCount++
//		}
//		checkArrayLength++
//		j++
//		if matchCount == sequenceLength{
//			return true
//		}
//
//		if checkArrayLength == arrayLength{
//			break
//		}
//	}
//	return false
//}

//
//func TwoNumberSum(array []int, target int) []int {
//	numberCheck := make(map[int]bool)
//	var ans []int
//	var y int
//	for i:=0; i <= len(array)-1 ; i++{
//		y = target - (array[i])
//		if numberCheck[y] == false{
//			numberCheck[array[i]]=true
//		}else{
//			ans = append(ans,array[i])
//			ans = append(ans,y)
//			return ans
//		}
//	}
//	return []int{}
//}

/*
 - sort array
 - first hour: 0-2 , 0-3

*/
//func largestTimeFromDigits(A []int) string {
//  ok := isValid(A)
//  var largestHour int
//  var largeHour int
//  var largestMin int
//  var largeMin int
//  if ok == true{
//  	for i:=0 ; i<= len(A)-1;i++{
//  		if A[i] == 2{
//  			largestHour = A[i]
//		}
//		if A[i] == 1{
//			if largestHour == 0{
//				largestHour = A[i]
//			}
//		}
//		if A[i] <= 3{
//			largeHour =A[i]
//		}
//
//		if A[i] == 5{
//			largestMin = A[i]
//		}else if A[i]<6{
//			largestMin =A[i]
//		}
//
//		if A[i] <10 {
//		}
//	  }
//  }else{
//  	return "nope"
//	}
//	myString := fmt.Sprintf("LargestHour: %v , LargeHour: %v , lgestMin: %v, lgmin:%v",largestHour, largeHour,largestMin,largeMin)
//	return myString
//}
//func isValid (arr []int)bool{
//	for i,_:=range arr{
//		if arr[i] ==2{
//			return true
//		}
//	}
//	return false
//}

//func TwoNumberSum(array []int, target int) []int {
//	var ans []int
//	var check int
//
//	for i:=0 ; i <= len(array) -1 ; i++{
//		for j:=1; j<=len(array-1); j++{
//			check=array[i] + array [j]
//
//			if check == target{
//				ans = append(ans, array[i])
//				ans = append(ans, array[j])
//			}
//		}
//	}
//	return ans
//}

//myList := linked_list.CreateList()
//myList.InsertAtHead("sugar")
//myList.Insert("water")
//myList.Insert("tomato")
//fmt.Println("Water index: ",myList.IndexOf("water"))
//myList.InsertAtPosition("beer",4)
//fmt.Println("What is at index 3: ", myList.PeekAtPosition( 3))
//fmt.Println("List Length: ",linked_list.ListLength(myList))
//linked_list.PrintList(myList)

//myStack := stack_queues.CreateStack()
//myStack.Add("Charlie")
//myStack.Add("Bravo")
//myStack.Add("Echo")
//myStack.Pop()
//
//myStack.PrintStack()
//fmt.Println(myStack.Search("Echo"))
//sorted := stack_queues.ReverseStack(myStack)
//sorted.Peek()

//values := []string{"paris", "shells", "carter", "brownie"}
//data := []string{"34", "30", "7", "6"}
//
//myTree := &bst.Tree{}
//for i:=0;i<len(values);i++{
//	err:= myTree.Insert(values[i], data[i])
//	if err != nil{
//		log.Fatal("Error", values[i], ":", err)
//	}
//}
//
//fmt.Print("Sort Vals: |")
//myTree.PreOrderTraverse(myTree.Root,func(n *bst.Node) {fmt.Print(n.Value, " : ", n.Data, " | ")})
//fmt.Println()
//
//fmt.Print("PostOrder Vals: |")
//myTree.PostOrderTraverse(myTree.Root,func(n *bst.Node) {fmt.Print(n.Value, " : ", n.Data, " | ")})
//fmt.Println()
//
//fmt.Print("InOrder Vals: |")
//myTree.InOrderTraverse(myTree.Root,func(n *bst.Node) {fmt.Print(n.Value, " : ", n.Data, " | ")})
//fmt.Println()
