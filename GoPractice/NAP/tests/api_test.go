package tests

import(
	"GoPractice/NAP"
	"net/http"
	"testing"
)
func TestAPICall(t *testing.T){
	api := NAP.NewAPI("https://httpbin.org")
	router := NAP.NewRouter()
	router.RegisterFunc(200,func(response *http.Response,_ interface{})error{
		return nil
	})
	res := NAP.NewResource()

}