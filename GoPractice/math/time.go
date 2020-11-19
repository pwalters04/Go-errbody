package main

import (
	"fmt"
	"time"
)


func main () {
	now := time.Now()
	layout := "20060102"
	fmt.Println(now.Date())
	inuput := now.Format(layout)
	fmt.Println(inuput)
	t, _ := time.Parse(layout,now.Format(layout) )
	fmt.Println(t)

}