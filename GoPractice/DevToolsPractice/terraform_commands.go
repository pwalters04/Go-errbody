package main

import (
	"fmt"
	"os"
)
func main(){
	cmdName := "terraform"
	cmdArg:= "plan"
	os.Chdir("~/develop/driveclutch/terraform/clutch-dev1/rds_grafanadb/")

	mydir, err := os.Getwd()
	if err == nil {
		fmt.Println(mydir)
		fmt.Println(cmdName + " " + cmdArg)
	} else{
		fmt.Println(err)
	}
}