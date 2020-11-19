package main
import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

const message = "Welcome Girlfriend! I am waiting for intellj to catch up"
func main (){
	mux:=http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","text/plain;  charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})
	err:= http.ListenAndServe(":8080",mux)
	if err != nil{
		log.Fatalf("Server Failed to Start: %v",err)
	}
}