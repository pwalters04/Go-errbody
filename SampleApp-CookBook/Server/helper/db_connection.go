package helper

import(
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"net/http"
)

// ErrorResponse : This is error model.
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func DBConnection() *mongo.Collection{
	 clientOpt := options.Client().ApplyURI( "mongodb+srv://newAdmin:test@cookbook.53s1p.mongodb.net/Cookbook?retryWrites=true&w=majority")
	 client,err := mongo.Connect(context.TODO(),clientOpt)
	 if err != nil{
	 	log.Fatal(err)
	 }
	fmt.Println("Connected To Database")

	 collection := client.Database("Cookbook").Collection("books")

	return collection
}

func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}