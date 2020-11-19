package tests

import(
	"GoPractice/NAP"
	"testing"
)
func TestEndpointTemplate(t *testing.T){
	res := &NAP.RestResource{
		Endpoint: "/user/{{.user}}",
		Method:   "GET",
		Router:   NAP.NewRouter(),
	}
	renderedEndpoint := res.RenderEndpoint(map[string]string{
		"user":"paris",
	})
	if renderedEndpoint != "/user/paris"{
		t.Fail()
	}
}