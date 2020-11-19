package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := http.DefaultClient
	//resp, err:=http.Get("https://httpbin.org/get?search=something") Get / Search Query
	//resp, err := http.Post("https://httpbin.org/post", "text/plain", strings.NewReader("This is the request content")) Post with payload
	req, err := http.NewRequest("GET", "http://httpbin.org/get", nil)
	if err != nil {
		log.Fatalln("Unable to create a request")
	}
	resp, err :=
		client.Do(req)
	if err != nil {
		log.Fatalln("Unable to get response")
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read content")
	}
	fmt.Println((string(content)))

}
