package linked_list

import (
	"fmt"
	"os"
)

type Node struct {
	Next *Node
	Data string
}

type List struct {
	Head *Node
	Tail *Node
	length int
}

func CreateList() *List{
	return &List{Head: nil,Tail: nil,length: 0}
}
func(list *List) InsertAtHead(data string) {
		dataInput :=&Node{
			Data: data,
		}

		if list.Head == nil{
			list.Head = dataInput
		} else {
			dataInput.Next =list.Head
			list.Head =dataInput
		}
		list.length ++

		return
}

func(list *List) InsertAtTail(data string){
		dataInput := &Node{Data: data,}
		if list.Tail == nil {
			list.Tail = dataInput
		}else{
			current := list.Head
			if current.Next != nil {
				current = current.Next
			}
			current.Next = dataInput
		}
		list.length ++
		return
}
func(list *List) Insert(data string){
	if list.Head == nil {
		list.InsertAtHead(data)
	}

	dataInput := &Node{
		Data: data,
	}
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next =dataInput
	list.length ++
	return
}

func (list *List) InsertAtPosition(data string, index int){
		if  index < 0{
			fmt.Println("Out Of Bounds. List Length: ",list.length)
			os.Exit(1)
		}

		if index == 0 {
			list.InsertAtHead(data)
			return
		}
		if index == list.length{
			list.InsertAtTail(data)
			return
		}
		dataInput := &Node{Data: data,}
		current := list.Head

		for j:=0 ; j < index-2 ;j++{
			current = current.Next
		}
		dataInput.Next = current.Next
		current.Next = dataInput

		list.length ++
		return
}

func (list *List) PeekAtPosition(index int) string{
	if index > list.length || index < 0{
		fmt.Println("Out Of Bounds. List Length: ",list.length)
		os.Exit(1)
	}
	current := list.Head
	if index == 0{
		return current.Data
	}
	i:=1
	for current.Next!=nil{
		if i == index{
			return current.Data
		}
		current = current.Next
		i++
	}

		return "Not Found"
}

func (list *List) IndexOf(Seekdata string) int{
	if list.Head == nil {
		fmt.Println("Empty List")
		os.Exit(1)
	}

	current := list.Head
	 i:=0
	for {
		if current.Data == Seekdata{
			return i
		}
		if current.Next == nil {
			fmt.Println("Not Found")
			return -1
		}
		current = current.Next
		i++
	}
}



func (list *List) SearchList(Seekdata string) (bool,int){
	 if list.Head == nil {
	 	fmt.Println("Empty List")
	 	os.Exit(1)
	 }

	current := list.Head
	for i:=0 ; i <= list.length ;i++{
		if current.Data == Seekdata{
			fmt.Println("Found")
			return true, list.IndexOf(Seekdata)
		}
			current = current.Next
	}

	 return false,-1
}
func ListLength (myList *List)int{
	return myList.length
}
func  PrintList(myList *List)error{
	if myList.Head == nil {
		fmt.Println("Empty List")
		os.Exit(1)
	}

	current := myList.Head
	fmt.Println(current.Data)
	for current.Next != nil{
		current = current.Next
		fmt.Println(current.Data)
	}
	return nil

}

func (list *List) RemoveAtHead() error { // TODO: Implement

	return nil
}