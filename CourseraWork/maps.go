package main

import (
	"fmt"
	"sort"
	"strings"
)

func main (){
	var personAge = map[string]int{
		"Rajeev": 25,
		"James":  32,
		"Sarah":  29,
	}

	_ = max(personAge)

}
func max(people map[string]int) string{
	result := map[int][]string{}
	var valueHolder []int
	var theKeys = make([]string, len(people))
	var i = 0

	for k,v := range people{
		result[v] =append(result[v],k)
	}
	for k := range result{
		valueHolder = append(valueHolder,k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(valueHolder)))
	for _,k := range valueHolder{
		for _,s := range result[k]{
			fmt.Println(s,k)
		}
	}

	for key,_ := range result{
		theKeys[i] = key
		i+=1
	}
	println(strings.Join(theKeys, ", "))
	println(len(theKeys))


	return ""
}