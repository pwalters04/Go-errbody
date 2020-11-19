package stack_queues

import (
	"fmt"
	"os"
	"strings"
)

type Stack []string

func CreateStack() *Stack{
	return &Stack{}
}

func(s *Stack) IsEmpty() bool{
	if len(*s)== 0{
		return true
	}
	return false
}

func(stack *Stack) Add(item string){
	*stack = append(*stack,item)
}

func(stack *Stack) Pop() string{
		if stack.IsEmpty() == true{
			fmt.Println("Empty Stack")
		}
		topIndex := len(*stack)-1
		fmt.Println("topIndex:",topIndex)
		item := (*stack)[topIndex]
		fmt.Println("item: ",item)
		*stack = (*stack)[:topIndex]
		fmt.Println("length: ",len(*stack))

		return item
}

func (s *Stack)Peek(){
	topIndex := len(*s)-1
	item := (*s)[topIndex]
	fmt.Println(item)

}

func(s *Stack) Search(seekItem string)bool{

	for _,key := range *s{
		if strings.Contains(key,seekItem) == true {
			return true
		}
	}
	return false
}

func (s *Stack) PrintStack(){
	topItem := len(*s)-1
	for i:= 0; i <= len(*s)-1 ;i++{
		fmt.Println((*s)[topItem])
		topItem --
	}
}

func ReverseStack(yourStack *Stack) *Stack{
	if yourStack.IsEmpty() == true {
		fmt.Println("Empty Stack")
		os.Exit(1)
	}

	tempStack := CreateStack()
	//for  yourStack{
	//	tempItem := yourStack.Pop()
	//	tempStack.Add(tempItem)
	//}
	return tempStack

}