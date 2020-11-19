package tests

import (
	NAP2 "GoPractice/NAP"
	"net/http"
	"testing"
)
func TestProcessRequest(t *testing.T){
	client:= NAP2.NewClient()
	router := NAP2.NewRouter()
	router.RegisterFunc(200,func(resp *http.Response,_ interface{})error{
		return nil
	})
	res:= NAP2.NewResource("/get", "GET",router)
	if err :=client.ProcessRequest("https://httpbin.org",res,nil); err != nil{
		t.Fail()
	}
}