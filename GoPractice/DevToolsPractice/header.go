package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://httpbin.org/basic-auth/user/passw0rd", nil)
	if err != nil {
		log.Fatalln("Unable to create response")
	}

	req.SetBasicAuth("user", "passw0rd")

	// ***** Long Way *********
	//buffer := &bytes.Buffer{}
	//enc := base64.NewEncoder(base64.URLEncoding,buffer)
	//enc.Write([]byte("user:passw0rd"))
	//enc.Close()
	//encodedCred, err :=buffer.ReadString('\n')
	//if err!= nil && err.Error()!="EOF"{
	//		log.Fatalln("Failed to read encoded buffer")
	//}
	//
	//req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedCred))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("Unable to read response")
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("unable to read content")
	}
	fmt.Println(string(content))
	fmt.Println("The Status Code:", resp.StatusCode)
}
