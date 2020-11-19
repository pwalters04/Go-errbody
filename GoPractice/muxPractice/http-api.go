package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func test(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This only a test"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", test)

	http.ListenAndServe(":5000", router)

}
