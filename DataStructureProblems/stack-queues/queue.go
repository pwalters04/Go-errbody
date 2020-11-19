package stack_queues

import "fmt"

type Data []int


func CreateQueue() *Data{
	return &Data{}
}
func(d *Data) Enqueue(number int ){
	*d = append(*d,number)
}

func(d *Data) Dequeue(){
	item :=(*d)[0]
	fmt.Println("Dequeue: ", item)
	*d = (*d)[1:]
}